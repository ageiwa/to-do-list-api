package user

import (
	"to-do-list-api/src/shared"
)

type User struct {
	Id int
	Email string
	Hash string
}

func CreateUser(user User) error {
	q := "INSERT INTO users (email, hash) VALUES (?, ?)"
	_, err := db.Conn.Exec(q, user.Email, user.Hash)

	return err
}

func FindUserByEmail(email string) (User, error) {
	q := "SELECT id, email, hash FROM users WHERE email = ?"
	user := User{}
	err := db.Conn.QueryRow(q, email).Scan(&user.Id, &user.Email, &user.Hash)

	if err != nil {
		return user, err
	}

	return user, nil
}