package db

import (
	"to-do-list-api/src/entities"
	"to-do-list-api/src/controllers"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepositoryMySQL struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) controllers.UserRespository {
	return &UserRepositoryMySQL{db: db}
}

func (r *UserRepositoryMySQL) CreateUser(user entities.User) error {
	q := "INSERT INTO users (email, hash) VALUES (?, ?)"
	_, err := r.db.Exec(q, user.Email, user.Hash)

	return err
}

func (r *UserRepositoryMySQL) FindUserByEmail(email string) (entities.User, error) {
	q := "SELECT id, email, hash FROM users WHERE email = ?"
	user := entities.User{}
	err := r.db.QueryRow(q, email).Scan(&user.Id, &user.Email, &user.Hash)

	if err != nil {
		return user, err
	}

	return user, nil
}