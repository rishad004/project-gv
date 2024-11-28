package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rishad004/project-gv/user-service/internal/domain"
	"github.com/spf13/viper"
)

func JwtCreate(userId int, Username string, Role string) (string, error) {
	claims := domain.Claims{
		UserID: uint(userId),
		Name:   Username,
		Role:   Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))

	return tokenString, err
}
