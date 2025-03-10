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

func CreateTask(task Task) error {
	q := "INSERT INTO tasks (title, desc) VALUES (?, ?)"
	_, err := db.Conn.Exec(q, task.Title, task.Desc)

	return err
}