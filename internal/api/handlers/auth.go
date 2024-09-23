package handlers

import (
    "encoding/json"
    "net/http"
    "packman-backend/internal/models"
    "packman-backend/internal/services"
)

// SignUpHandler handles user registration
func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var creds models.Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Call the service to register the user
    err = services.RegisterUser(creds.Username, creds.Password)
    if err != nil {
        if err == services.ErrUserExists {
            http.Error(w, "User already exists", http.StatusConflict)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds models.Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Call the service to authenticate the user and generate a JWT token
    token, err := services.AuthenticateUser(creds.Username, creds.Password)
    if err != nil {
        if err == services.ErrInvalidCredentials {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}
