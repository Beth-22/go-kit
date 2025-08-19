package handlers

import (
	"JWT-GO3/internal/models"
	"JWT-GO3/internal/services"
	"JWT-GO3/internal/storage"
	"encoding/json"
	"net/http"
)

//signup
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	 // Default role to "user" if empty
    if user.Role == "" {
        user.Role = "user"
    }

	if err := storage.CreateUser(user); 
	err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created successfully!"))
}


//login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	user, err := storage.GetUser(creds.Username)
	if err != nil || user.Password != creds.Password {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

token, err := services.GenerateJWT(user.Username, user.Role)
	if err != nil {
		http.Error(w, "token generation failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))//response to the login request
}
