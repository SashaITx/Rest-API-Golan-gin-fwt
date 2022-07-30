// этот пакет общается с БД и связан с пакетом сервисов

package repository

import (
	rest "Rest_API_Golan-gin-fwt"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GetUser(username, password string) (rest.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
