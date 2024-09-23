package models

import "github.com/dgrijalva/jwt-go"

// Claims defines the structure for JWT claims
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}
