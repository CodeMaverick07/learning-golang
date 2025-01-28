package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error) {
	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print(err)
		return "",err 
	}
	return string(hashedPassword),nil
} 

func ComparePasswords(password string,hashedPassword string) (bool) {
	
	err:= bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	
	return err == nil
}