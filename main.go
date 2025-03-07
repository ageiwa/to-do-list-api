package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"to-do-list-api/src/handlers"
	"to-do-list-api/src/shared"
)

func main() {
	godotenv.Load()

	db.NewDB()

	http.HandleFunc("/register", user.Register)
	http.HandleFunc("/login", user.Login)

	println("Server is started...")
	http.ListenAndServe(":8080", nil)
}
