package handlers

import (
	"to-do-list-api/src/entities"

	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterOptions struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func errResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	resp := make(map[string]any)
	resp["message"] = message

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func successResponse(w http.ResponseWriter, resp map[string]any, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	
	opt := RegisterOptions{}

	if err := decoder.Decode(&opt); err != nil {
		errResponse(w, "Error decode " + err.Error(), http.StatusBadRequest)
		return
	}

	_, err := entities.FindUserByEmail(opt.Email)

	if err == nil {
		errResponse(w, "User already exists", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(opt.Password), bcrypt.DefaultCost)

	if err != nil {
		errResponse(w, "Error hash password", http.StatusInternalServerError)
		return
	}

	err = entities.CreateUser(entities.User{
		Email: opt.Email,
		Hash: string(hash),
	})
	
	if err != nil {
		log.Fatal(err.Error())
	}

	successResponse(w, map[string]any{
		"success": true,
	}, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	opt := RegisterOptions{}
	
	if err := decoder.Decode(&opt); err != nil {
		errResponse(w, "Error decode " + err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := entities.FindUserByEmail(opt.Email)

	if err != nil {
		errResponse(w, "User not found", http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(opt.Password)); err != nil {
		errResponse(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	payload := jwt.MapClaims{
		"id": user.Id,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	
	t, err := token.SignedString([]byte("my-super-sign"))

	if err != nil {
		errResponse(w, "JWT token sign " + err.Error(), http.StatusInternalServerError)
		return
	}

	successResponse(w, map[string]any{
		"token": t,
	}, http.StatusOK)
}