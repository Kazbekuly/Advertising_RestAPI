package repository

import (
	"Advertising/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, login, password) values ($1,$2,$3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
