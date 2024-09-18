package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func main() {
	// เชื่อมต่อฐานข้อมูล
	dsn := "root:root@tcp(localhost:3306)/auth_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// ตรวจสอบว่าเป็น POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			log.Printf("Error decoding request body: %v", err)
			return
		}

		// ตรวจสอบข้อมูลผู้ใช้ในฐานข้อมูล
		var storedPassword, email, role string
		err = db.QueryRow("SELECT email, password, role FROM users WHERE username = ?", creds.Username).Scan(&email, &storedPassword, &role)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				log.Printf("User not found: %v", err)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error querying the database: %v", err)
			return
		}

		// เปรียบเทียบรหัสผ่าน
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password))
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			log.Printf("Password comparison failed: %v", err)
			return
		}

		// สร้าง JWT token
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &Claims{
			Username: creds.Username,
			Email:    email,
			Role:     role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error signing the token: %v", err)
			return
		}

		// ส่งข้อมูลกลับไปยัง client
		response := map[string]string{
			"token":    tokenString,
			"username": creds.Username,
			"email":    email,
			"role":     role,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error encoding response: %v", err)
		}
	})

	fmt.Println("Server is running on port 7777")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
