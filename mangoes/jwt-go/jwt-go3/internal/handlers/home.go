package handlers

import (
    "fmt"
    "net/http"

    "JWT-GO3/internal/middleware"
    "JWT-GO3/internal/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    claims, ok := r.Context().Value(middleware.UserContextKey).(*services.Claims)
    if !ok || claims == nil {
        http.Error(w, "User not found in context", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Welcome, %s! You are logged in as %s.", claims.Username, claims.Role)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
    claims, ok := r.Context().Value(middleware.UserContextKey).(*services.Claims)
    if !ok || claims == nil {
        http.Error(w, "User not found in context", http.StatusInternalServerError)
        return
    }

    if claims.Role != "admin" {
        http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
        return
    }

    fmt.Fprintf(w, "Hello, Admin %s!", claims.Username)
}
