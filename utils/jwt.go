package utils

import (
	"errors"
	"fmt"
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

func VerifyToken(token string) (int64, error) {
    parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
        // Check if the signing method is what we expect
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })

    if err != nil {
        return 0, fmt.Errorf("failed to parse token: %w", err)
    }

    if !parsedToken.Valid {
        return 0, errors.New("token is invalid")
    }

    claims, ok := parsedToken.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid token claims")
    }
 
    // JWT numbers are typically float64, so we need to convert
    userIdFloat, ok := claims["userId"].(float64)
    if !ok {
        return 0, errors.New("invalid userId claim")
    }
    
    // Convert float64 to int64
    userId := int64(userIdFloat)
    fmt.Print(userId)
    
    return userId, nil
}