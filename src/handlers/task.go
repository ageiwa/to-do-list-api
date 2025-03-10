package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"to-do-list-api/src/entities"

	"github.com/golang-jwt/jwt/v5"
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
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		errResponse(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		errResponse(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
		return
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my-super-sign"), nil
	})

	if err != nil || !token.Valid {
		errResponse(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		errResponse(w, "Invalid claims", http.StatusInternalServerError)
		return
	}

	exp, ok := claims["exp"].(float64)

	if !ok {
		errResponse(w, "Wrong exp type", http.StatusUnauthorized)
		return
	}

	if int64(exp) < time.Now().Unix() {
		errResponse(w, "Token expired", http.StatusUnauthorized)
		return
	}

	id, ok := claims["id"].(string)

	if !ok {
		errResponse(w, "Wrong id type", http.StatusUnauthorized)
		return 
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	opt := CreateTaskOptions{}

	if err := decoder.Decode(&opt); err != nil {
		errResponse(w, "Error decode: " + err.Error(), http.StatusBadRequest)
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		panic(err.Error())
	}

	if err := entities.CreateTask(opt.Title, opt.Desc, int(idInt)); err != nil {
		errResponse(w, "Cant create task: " + err.Error(), http.StatusBadRequest)
		return
	}

	println("create task")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	println("get tasks")
}