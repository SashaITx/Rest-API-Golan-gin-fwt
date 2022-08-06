// этот пакет общается с БД и связан с пакетом сервисов

package repository

import (
	rest "Rest_API_Golan-gin-fwt"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GetUser(username, password string) (rest.User, error)
}

type Repository struct {
	Authorization
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user rest.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (rest.User, error) {
	var user rest.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
