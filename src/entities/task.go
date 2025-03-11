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

func CreateTask(title string, desc string, userId int) (int, error) {
	q := "INSERT INTO tasks (title, description, userId) VALUES (?, ?, ?)"
	res, err := db.Conn.Exec(q, title, desc, userId)

	if err != nil {
		return 0, err
	}
	
	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}