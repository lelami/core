package repository

import (
	"core"
	"github.com/jmoiron/sqlx"
	"github.com/roistat/go-clickhouse"
)

type Authorization interface {
	CreateAuthUser(user core.AuthUser) (int, error)
	GetAuthUser(phone int64) (core.AuthUser, error)
	DeleteAuthUser(phone int64) (int64, error)
}

type Registration interface {
	CreateUser(user core.User) (int, error)
	GetUser(phone int64) (core.User, error)
}

type Logging interface {
	WriteLog(log string) error
}

type Repository struct {
	Authorization
	Registration
	Logging
}

func NewRepository(db *sqlx.DB, CH *clickhouse.Conn) *Repository {
	return &Repository{
		Registration:  NewUsersPostgres(db),
		Authorization: NewAuthPostgres(db),
		Logging:       NewLogClickHouse(CH),
	}
}
