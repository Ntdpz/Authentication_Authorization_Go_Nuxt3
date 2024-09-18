package middleware

import (
	"github.com/rs/cors"
)

// NewCORSHandler สร้าง CORS middleware
func NewCORSHandler() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})
}
