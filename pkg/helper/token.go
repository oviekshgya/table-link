package helper

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string) (string, error) {
	claims := &Claims{
		UserID:   id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("supersecretkey"))
}

func ValidateToken(tokenString string) (bool, uint, string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("supersecretkey"), nil
	})

	if err != nil {
		return false, 0, "", err
	}

	if !token.Valid {
		return false, 0, "", nil
	}

	return true, claims.UserID, claims.Username, nil
}
