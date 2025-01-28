package routes

import (
	"fmt"

	"example.com/api/models"
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)
 

func signup(contex *gin.Context){
	var user models.User
	 
	err:= contex.ShouldBindJSON(&user)
	if err != nil {
       fmt.Println(err)
	   contex.JSON(400, gin.H{"error": "Invalid request while sign in user"})
	   return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		contex.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	contex.JSON(200, gin.H{"message": "User created successfully"})
} 

func login(contex *gin.Context){
	var user models.User
	 
	err:= contex.ShouldBindJSON(&user)
	if err != nil {
	   fmt.Println(err)
	   contex.JSON(400, gin.H{"error": "Invalid request while sign in user"})
	   return
	}

	err = user.Authenticate()
	if err != nil {
		fmt.Println(err)
		contex.JSON(500, gin.H{"error": "Failed to authenticate user"})
		return
	}
	token,err:= utils.GenerateToken(user.Email,user.ID)
	if err != nil {
		fmt.Println(err)
		contex.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	contex.JSON(200, gin.H{"message": "User authenticated successfully","token":token})
	contex.SetCookie("token",token,3600,"/","localhost",false,true)
} 