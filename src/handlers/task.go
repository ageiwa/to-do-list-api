package handlers

import (
	// "to-do-list-api/src/entities"
	"encoding/json"
	"net/http"
)

type CreateTaskOptions struct {
	Title string
	Desc string
}

func TaskHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		CreateTask(w, r)
	} else if r.Method == http.MethodGet {
		GetTasks(w, r)
	} else {
		errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	opt := CreateTaskOptions{}

	if err := decoder.Decode(&opt); err != nil {
		errResponse(w, "Error decode " + err.Error(), http.StatusBadRequest)
		return
	}

	// entities.CreateTask()
	println("create task")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	println("get tasks")
}