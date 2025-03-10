package entities

import (
	"time"
	"to-do-list-api/src/shared"
)

type Task struct {
	Id int
	Title string
	Desc string
	CreatedAt time.Time
}

func CreateTask(title string, desc string, userId int) error {
	q := "INSERT INTO tasks (title, description, userId) VALUES (?, ?, ?)"
	_, err := db.Conn.Exec(q, title, desc, userId)

	return err
}