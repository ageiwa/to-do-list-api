package handlers

import (
	"encoding/json"
	"net/http"
	"to-do-list-api/src/entities"
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
	userIdFloat, ok := r.Context().Value(userIdKey).(float64)

	if !ok {
		errResponse(w, "User id not found", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	opt := CreateTaskOptions{}

	if err := decoder.Decode(&opt); err != nil {
		errResponse(w, "Error decode: " + err.Error(), http.StatusBadRequest)
		return
	}

	taskId, err := entities.CreateTask(opt.Title, opt.Desc, int(userIdFloat))

	if err != nil {
		errResponse(w, "Cant create task: " + err.Error(), http.StatusBadRequest)
		return
	}

	successResponse(w, map[string]any{
		"taskId": taskId,
	}, http.StatusOK)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	userIdFloat, ok := r.Context().Value(userIdKey).(float64)

	if !ok {
		errResponse(w, "User id not found", http.StatusInternalServerError)
		return
	}

	userId := int(userIdFloat)
	tasks, err := entities.GetTasks(userId)

	if err != nil {
		errResponse(w, "Cant get tasks", http.StatusBadRequest)
		return
	}

	successResponse(w, map[string]any{
		"tasks": tasks,
	}, http.StatusOK)
}