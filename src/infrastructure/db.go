package db

import (
	"to-do-list-api/src/entities"

	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func NewDB() *sql.DB {
	USER, _ := os.LookupEnv("MYSQL_USER")
	PASSWORD, _ := os.LookupEnv("MYSQL_PASSWORD")
	PORT, _ := os.LookupEnv("MYSQL_PORT")

	src := USER + ":" + PASSWORD + "@(127.0.0.1:" + PORT + ")/mydb?parseTime=true"
	conn, _ := sql.Open("mysql", src)

	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}

	return conn
}

func CreateUser(user entities.User) error {
	q := "INSERT INTO users (email, hash) VALUES (?, ?)"
	_, err := conn.Exec(q, user.Email, user.Hash)

	return err
}

func FindUserByEmail(email string) (entities.User, error) {
	q := "SELECT id, email, hash FROM users WHERE email = ?"
	user := entities.User{}
	err := conn.QueryRow(q, email).Scan(&user.Id, &user.Email, &user.Hash)

	if err != nil {
		return user, err
	}

	return user, nil
}