package repository

import (
	"github.com/DavelPurov777/microblog/services/api/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query, args, err := psql.
		Insert(usersTable).
		Columns("name", "username", "password_hash").
		Values(user.Name, user.Username, user.Password).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
