package controllers

import (
	// "to-do-list-api/entities"

	"net/http"
	"encoding/json"
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
	// if r.Method != http.MethodPost {
	// 	errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// decoder := json.NewDecoder(r.Body)
	// decoder.DisallowUnknownFields()
	
	// opt := RegisterOptions{}

	// if err := decoder.Decode(&opt); err != nil {
	// 	errResponse(w, "Error decode " + err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// _, err := db.FindUserByEmail(opt.Email)

	// if err == nil {
	// 	errResponse(w, "User already exists", http.StatusBadRequest)
	// 	return
	// }

	// hash, err := bcrypt.GenerateFromPassword([]byte(opt.Password), bcrypt.DefaultCost)

	// if err != nil {
	// 	errResponse(w, "Error hash password", http.StatusInternalServerError)
	// 	return
	// }

	// err = db.CreateUser(entities.User{
	// 	Email: opt.Email,
	// 	Hash: string(hash),
	// })
	
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// successResponse(w, map[string]any{
	// 	"success": true,
	// }, http.StatusOK)
}