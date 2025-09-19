package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type Claims struct {
	Email string `json:"username"`
	jwt.RegisteredClaims
}

func SignToken(email string) string {
	jwt_secret := os.Getenv("JWT_SECRET")
	claims := &Claims{
		Email: email ,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // 1 hour expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-jwt-auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString , err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}


func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return secret key as []byte
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println("Token parse error:", err)
		return nil,err
	}

	// Check if token is valid
	if !token.Valid {
		log.Println("Invalid token")
		return nil , fmt.Errorf("token is invalid")
	}
	
	// Token is valid
	return claims, nil
}
