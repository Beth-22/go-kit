package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"JWT-GO3/internal/handlers"
	"JWT-GO3/internal/middleware"
	"JWT-GO3/internal/storage"
)

func main() {
    // Connect to PostgreSQL using pgx (no other logic changed)
    err := storage.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to DB: ", err)
    }

    r := mux.NewRouter()

    // Public routes
    r.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

    // Protected routes
    api := r.PathPrefix("/api").Subrouter() //This creates a subrouter called api that will handle all routes starting with /api.

    api.Use(middleware.AuthMiddleware)
    //This means every request to any route under /api must pass through this middleware first.The AuthMiddleware verifies the JWT token from the request’s Authorization header.


    
    api.HandleFunc("/home", handlers.HomeHandler).Methods("GET")
    api.HandleFunc("/admin", handlers.AdminHandler).Methods("GET")

    log.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
