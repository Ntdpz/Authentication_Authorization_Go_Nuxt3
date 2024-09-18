package routes

import (
	"auth-api/controllers"
	"database/sql"
	"net/http"
)

// SetupRoutes สร้างและตั้งค่า routes
func SetupRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", controllers.RegisterHandler(db))
	mux.HandleFunc("/login", controllers.LoginHandler(db))
	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged out successfully"))
	})
	// New test route
	mux.HandleFunc("/test", controllers.TestHandler)
	return mux
}
