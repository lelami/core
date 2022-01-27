package repository

import (
	"core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) CreateUser(user core.User) (int, error) {
	var id int

	founded, err := r.GetUser(user.Phone)
	if err != nil {
		query := fmt.Sprintf("INSERT INTO %s (name, phone) values ($1, $2) RETURNING id", usersTable)

		row := r.db.QueryRow(query, user.Name, user.Phone)
		if err = row.Scan(&id); err != nil {
			return 0, err
		}
		return id, nil
	}
	return founded.Id, nil
}

func (r *UsersPostgres) GetUser(phone int64) (core.User, error) {
	var user core.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE phone=$1", usersTable)
	err := r.db.Get(&user, query, phone)

	return user, err
}
