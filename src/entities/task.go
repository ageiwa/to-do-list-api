package entities

import (
	"strings"
	"to-do-list-api/src/shared"
)

type Task struct {
	Id int
	Title string
	Desc string
	CreatedAt string
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

func GetTasks(userId int) ([]Task, error) {
	q := "SELECT id, title, description, createdAt FROM tasks WHERE userId = ?"
	rows, err := db.Conn.Query(q, userId)

	if err != nil {
		return []Task{}, err
	}

	defer rows.Close()

	tasks := []Task{}

	for rows.Next() {
		task := Task{}

		if err := rows.Scan(&task.Id, &task.Title, &task.Desc, &task.CreatedAt); err != nil {
			return []Task{}, err
		}

		task.CreatedAt = strings.Split(task.CreatedAt, "T")[0]
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}