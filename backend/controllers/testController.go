// backend/controllers/testController.go
package controllers

import (
	"log"
	"net/http"
)

// TestHandler handles requests to the /test endpoint
func TestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"message": "This is a test endpoint!"}`)); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
	}
}
