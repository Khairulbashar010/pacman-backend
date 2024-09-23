package services

import (
    "errors"
    "packman-backend/internal/models"
    "packman-backend/utils"
)

// Custom error messages
var (
    ErrInvalidCredentials = errors.New("invalid username or password")
    ErrUserExists         = errors.New("user already exists")
)

// RegisterUser registers a new user by hashing the password and saving it to the database
func RegisterUser(username, password string) error {
    // Check if the user already exists
    existingUser, _ := models.GetUserByUsername(username)
    if existingUser != nil {
        return ErrUserExists
    }

    // Create a new user with the hashed password
    user := models.User{
        Username: username,
        Password: password,
    }

    err := user.CreateUser()
    if err != nil {
        return err
    }

    return nil
}

// AuthenticateUser authenticates the user by comparing their credentials and returns a JWT token
func AuthenticateUser(username, password string) (string, error) {
    // Fetch the user by username
    user, err := models.GetUserByUsername(username)
    if err != nil || user == nil {
        return "", ErrInvalidCredentials
    }

    // Compare the provided password with the hashed password in the database
    if !user.ComparePassword(password) {
        return "", ErrInvalidCredentials
    }

    // Generate a new JWT token using utils/jwt.go
    token, err := utils.GenerateJWT(user.Username)
    if err != nil {
        return "", err
    }

    return token, nil
}
