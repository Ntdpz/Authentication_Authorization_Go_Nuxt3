// backend/config/config.go
package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB เชื่อมต่อฐานข้อมูล
func ConnectDB() (*sql.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/auth_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	// ทดสอบการเชื่อมต่อฐานข้อมูล
	if err := db.Ping(); err != nil {
		log.Printf("Error connecting to the database: %v", err)
		return nil, err
	}

	return db, nil
}
