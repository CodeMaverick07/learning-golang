package main

import (
	"net/http"

	"example.com/api/db"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)
	


func main() {
	db.InitDB()
	server:=gin.Default()
	server.GET("/events",getEvents)
	server.POST("/events",createEvents)
	server.Run(":8080")
}

func getEvents(contex *gin.Context){
	events := models.GetAllEvents()
	contex.JSON(http.StatusOK,events)
}

func createEvents (contex *gin.Context){
	var event models.Event
	err:=contex.ShouldBindJSON(&event)
	if err != nil {
		contex.JSON(http.StatusBadRequest,gin.H{"message":"could not the requried fields"})
		return
	}
	event.ID =1 
	event.UserID = "1"
	event.Save()
	contex.JSON(http.StatusCreated,gin.H{"message":"message created","event":event})
 }