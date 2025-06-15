package main

import (
	"context"
	"log"
	"net/http"
	"user-management-service/db"
	"user-management-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	defer db.Conn.Close(context.Background())

	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
