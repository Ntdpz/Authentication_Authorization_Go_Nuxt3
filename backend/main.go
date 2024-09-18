package main

import (
	"auth-api/config"
	"auth-api/middleware"
	"auth-api/routes"
	"log"
	"net/http"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// สร้าง CORS middleware
	corsHandler := middleware.NewCORSHandler()

	// สร้าง router
	r := routes.SetupRoutes(db)

	// ใช้งาน CORS middleware
	server := corsHandler.Handler(r)

	log.Println("Server is running on port 7777")
	log.Fatal(http.ListenAndServe(":7777", server))
}
