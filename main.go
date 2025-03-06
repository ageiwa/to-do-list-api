package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
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

var conn *sql.DB

func errResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	resp := make(map[string]any)
	resp["message"] = message

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func successResopnse(w http.ResponseWriter, resp map[string]any, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func register(w http.ResponseWriter, r *http.Request) {
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

	q := "SELECT * FROM users WHERE email = ?"
	rows, err := conn.Query(q, opt.Email)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.id, &user.email, &user.hash)
		
		if err != nil {
			log.Fatal(err.Error())
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	if len(users) > 0 {
		errResponse(w, "User already exists", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(opt.Password), bcrypt.DefaultCost)

	if err != nil {
		errResponse(w, "Error hash password", http.StatusInternalServerError)
		return
	}

	q = "INSERT INTO users (email, hash) VALUES (?, ?)"
	
	if _, err := conn.Exec(q, opt.Email, hash); err != nil {
		log.Fatal(err.Error())
	}

	successResopnse(w, map[string]any{
		"success": true,
	}, http.StatusOK)
}

func login(w http.ResponseWriter, r *http.Request) {
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

	q := "SELECT id, email, hash FROM users WHERE email = ?"
	user := User{}

	if err := conn.QueryRow(q, opt.Email).Scan(&user.id, &user.email, &user.hash); err != nil {
		if err == sql.ErrNoRows {
			errResponse(w, "User not found", http.StatusBadRequest)
			return
		} else {
			log.Fatal(err.Error())
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.hash), []byte(opt.Password)); err != nil {
		errResponse(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	successResopnse(w, map[string]any{
		"id": user.id,
		"email": user.email,
	}, http.StatusOK)
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	USER, _ := os.LookupEnv("MYSQL_USER")
	PASSWORD, _ := os.LookupEnv("MYSQL_PASSWORD")
	PORT, _ := os.LookupEnv("MYSQL_PORT")

	src := USER + ":" + PASSWORD + "@(127.0.0.1:" + PORT + ")/mydb?parseTime=true"
	conn, err := sql.Open("mysql", src)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	println("Server is started...")
	http.ListenAndServe(":8080", nil)
}