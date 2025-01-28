package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "secret" 

func GenerateToken(email string, userId int64) (string,error) {
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId":userId,
			"email":email,
			"exp":time.Now().Add(time.Hour * 24).Unix(),
			}) 
			return token.SignedString([]byte(secret)) 
} 