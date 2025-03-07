package db

import (
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func NewDB() {
	USER, _ := os.LookupEnv("MYSQL_USER")
	PASSWORD, _ := os.LookupEnv("MYSQL_PASSWORD")
	PORT, _ := os.LookupEnv("MYSQL_PORT")

	src := USER + ":" + PASSWORD + "@(127.0.0.1:" + PORT + ")/mydb?parseTime=true"
	Conn, _ = sql.Open("mysql", src)

	if err := Conn.Ping(); err != nil {
		panic(err.Error())
	}
}