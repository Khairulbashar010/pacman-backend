package utils

import (
    "errors"
    "time"
    "packman-backend/internal/config"
    "packman-backend/internal/models"

    "github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a new JWT token for the user
func GenerateJWT(username string) (string, error) {
    cfg := config.GetConfig()

    // Set token expiration time (e.g., 24 hours)
    expirationTime := time.Now().Add(24 * time.Hour)

    // Create the JWT claims, which include the username and expiry time
    claims := &models.Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Create the token using the claims and sign it with the secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims if it's valid
func ValidateToken(tokenString string) (*models.Claims, error) {
    cfg := config.GetConfig()

    // Parse the token
    claims := &models.Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Ensure the signing method is HMAC (HS256)
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return []byte(cfg.JWTSecret), nil
    })

    if err != nil {
        // Provide better error messages for debugging
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
