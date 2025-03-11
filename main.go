package main

import (
	"net/http"
	"to-do-list-api/src/handlers"
	"to-do-list-api/src/shared"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.NewDB()

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc("/tasks", handlers.TaskHanlder)

	println("Server is started...")
	http.ListenAndServe(":8080", nil)
}
