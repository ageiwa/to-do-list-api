package main

import (
	"to-do-list-api/src/entities"
	"to-do-list-api/src/infrastructure"
	"to-do-list-api/src/usecases"

	"database/sql"
	// "encoding/json"
	// "log"
	// "net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	// "golang.org/x/crypto/bcrypt"
)

type RegisterOptions struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// func errResponse(w http.ResponseWriter, message string, httpStatusCode int) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(httpStatusCode)

// 	resp := make(map[string]any)
// 	resp["message"] = message

// 	jsonResp, _ := json.Marshal(resp)
// 	w.Write(jsonResp)
// }

// func successResponse(w http.ResponseWriter, resp map[string]any, httpStatusCode int) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(httpStatusCode)

// 	jsonResp, _ := json.Marshal(resp)
// 	w.Write(jsonResp)
// }

// func register(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	decoder.DisallowUnknownFields()
	
// 	opt := RegisterOptions{}

// 	if err := decoder.Decode(&opt); err != nil {
// 		errResponse(w, "Error decode " + err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	_, err := db.FindUserByEmail(opt.Email)

// 	if err == nil {
// 		errResponse(w, "User already exists", http.StatusBadRequest)
// 		return
// 	}

// 	hash, err := bcrypt.GenerateFromPassword([]byte(opt.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		errResponse(w, "Error hash password", http.StatusInternalServerError)
// 		return
// 	}

// 	err = db.CreateUser(entities.User{
// 		Email: opt.Email,
// 		Hash: string(hash),
// 	})
	
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	successResponse(w, map[string]any{
// 		"success": true,
// 	}, http.StatusOK)
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		errResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	decoder.DisallowUnknownFields()

// 	opt := RegisterOptions{}
	
// 	if err := decoder.Decode(&opt); err != nil {
// 		errResponse(w, "Error decode " + err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	user, err := db.FindUserByEmail(opt.Email)

// 	if err != nil {
// 		errResponse(w, "User not found", http.StatusBadRequest)
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(opt.Password)); err != nil {
// 		errResponse(w, "Invalid email or password", http.StatusBadRequest)
// 		return
// 	}

// 	successResponse(w, map[string]any{
// 		"id": user.Id,
// 		"email": user.Email,
// 	}, http.StatusOK)
// }

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	USER, _ := os.LookupEnv("MYSQL_USER")
	PASSWORD, _ := os.LookupEnv("MYSQL_PASSWORD")
	PORT, _ := os.LookupEnv("MYSQL_PORT")

	src := USER + ":" + PASSWORD + "@(127.0.0.1:" + PORT + ")/mydb?parseTime=true"
	conn, _ := sql.Open("mysql", src)

	userRepo := db.NewUserRepository(conn)
	userUseCases := usecases.NewUserUseCase(userRepo)

	userUseCases.CreateUser(entities.User{
		Email: "super-puper",
		Hash: "31312312123",
	})

	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}

	// http.HandleFunc("/register", register)
	// http.HandleFunc("/login", login)

	println("Server is started...")
	// http.ListenAndServe(":8080", nil)
}