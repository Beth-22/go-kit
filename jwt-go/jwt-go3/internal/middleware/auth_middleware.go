package middleware

import (
    "net/http"
    "strings"
    "context"

    "JWT-GO3/internal/services"
)

type contextKey string

const (
    UserContextKey = contextKey("user")
)

// AuthMiddleware verifies the JWT token and adds user claims to context
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Missing or malformed token", http.StatusUnauthorized)
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        claims, err := services.VerifyJWT(tokenStr)
        if err != nil {
            http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
            return
        }

        // Add claims to context.This allows downstream handlers to access authenticated user info via context.


        ctx := context.WithValue(r.Context(), UserContextKey, claims)
        //Passes the request to the next handler in the chain.Replaces the original request’s context with the new one containing user claims.This effectively injects user info into the request flow.
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
