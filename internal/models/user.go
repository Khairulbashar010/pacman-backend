package models

import (
    "database/sql"
    "time"
    "packman-backend/internal/db"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"-"` // Exclude password from JSON output
    CreatedAt time.Time `json:"created_at"`
}

// CreateUser creates a new user in the database
func (u *User) CreateUser() error {
    hashedPassword, err := hashPassword(u.Password)
    if err != nil {
        return err
    }

    query := "INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)"
    _, err = db.DB.Exec(query, u.Username, hashedPassword, time.Now())
    if err != nil {
        return err
    }
    return nil
}

// GetUserByUsername fetches a user by username
func GetUserByUsername(username string) (*User, error) {
    var user User
    query := "SELECT id, username, password, created_at FROM users WHERE username = ?"
    err := db.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
    if err == sql.ErrNoRows {
        return nil, nil // No user found
    } else if err != nil {
        return nil, err
    }
    return &user, nil
}

// ComparePassword compares the hashed password with the provided plain password
func (u *User) ComparePassword(plainPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
    return err == nil
}

// hashPassword hashes a plain password using bcrypt
func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}
