package jwt

import (
	"fmt"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// GenerateToken creates a new JWT for a specific user
func GenerateToken(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

  res := GenerateToken("exampleUser")
  fmt.println(res)
  