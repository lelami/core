package repository

import (
	"core"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateAuthUser(user core.AuthUser) (int, error) {

	var id int
	query := fmt.Sprintf("INSERT INTO %s (phone, code) values ($1, $2) ON CONFLICT (phone) DO UPDATE SET code=$2 RETURNING id", authTable)

	row := r.db.QueryRow(query, user.Phone, user.Code)
	if err := row.Scan(&id); err != nil {
		return id, nil
	}

	go func() {
		t := time.NewTimer(60 * time.Second)
		<-t.C
		r.DeleteAuthUser(user.Phone)
	}()

	return id, nil
}

func (r *AuthPostgres) GetAuthUser(phone int64) (core.AuthUser, error) {
	var user core.AuthUser
	query := fmt.Sprintf("SELECT id, phone, code FROM %s WHERE phone=$1", authTable)
	err := r.db.Get(&user, query, phone)

	r.DeleteAuthUser(phone)

	return user, err
}

func (r *AuthPostgres) DeleteAuthUser(phone int64) (int64, error) {

	_, err := r.db.Exec(fmt.Sprintf("DELETE FROM %s where phone=$1", authTable), phone)

	return phone, err
}
