package middlewares

import (
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(contex *gin.Context) {
	token:= contex.Request.Header.Get("Authorization")

	if token == "" {
		contex.JSON(401,gin.H{"error":"No token found"})
		contex.Abort()
		return
	}
	userId,err := utils.VerifyToken(token)
	if err != nil {
		contex.JSON(401,gin.H{"error":"Invalid token"})
		contex.Abort()
		return
	}
	contex.Set("userId",userId) 
	contex.Next()

}