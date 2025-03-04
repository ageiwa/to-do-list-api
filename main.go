package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id int
	email string
	hash string
}

type RegisterOptions struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var idCounter = 0
var users []User

func errResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)

	resp := make(map[string]any)
	resp["message"] = message

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	
	opt := RegisterOptions{}
	err := decoder.Decode(&opt)

	if err != nil {
		errResponse(w, "Error decode " + err.Error(), http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.email == opt.Email {
			errResponse(w, "User already exist", http.StatusBadRequest)
			return
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(opt.Password), bcrypt.DefaultCost)

	if err != nil {
		errResponse(w, "Error hash password", http.StatusInternalServerError)
		return
	}

	idCounter++

	users = append(users, User{
		id: idCounter,
		email: opt.Email,
		hash: string(hash),
	})

	w.Write([]byte(`{ "success": true }`))
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	opt := RegisterOptions{}
	err := decoder.Decode(&opt)

	if err != nil {
		errResponse(w, "Error decode " + err.Error(), http.StatusBadRequest)
		return
	}

	isFind := false
	findUser := User{}

	for _, user := range users {
		if user.email == opt.Email {
			findUser = user
			isFind = true
			break
		}
	}

	if !isFind {
		errResponse(w, "User not found", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.hash), []byte(opt.Password))

	if err != nil {
		errResponse(w, "Invalid Email or password", http.StatusBadRequest)
		return
	}

	resp := make(map[string]any)
	resp["id"] = findUser.id
	resp["email"] = findUser.email

	jsonResp, _ := json.Marshal(resp)

	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	println("Server is started...")
	http.ListenAndServe(":8080", nil)
}