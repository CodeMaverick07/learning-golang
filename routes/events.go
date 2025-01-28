package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)


func getAllEvents(contex *gin.Context){
	events,err := models.GetAllEvents()
	if err!= nil {
		contex.JSON(http.StatusBadRequest,gin.H{"message":"not able to fetch data"})
		return 
	}
	contex.JSON(http.StatusOK,events)
}

func createEvents (contex *gin.Context){
	var event models.Event
	err:=contex.ShouldBindJSON(&event)
	if err != nil {
		contex.JSON(http.StatusBadRequest,gin.H{"message":"could not the requried fields"})
		return
	}
	event.UserID = "1"
	err = event.Save()
	if err != nil {
		contex.JSON(http.StatusBadRequest,gin.H{"message":"not able to save data"})
		return 
	}
	contex.JSON(http.StatusCreated,gin.H{"message":"message created","event":event})
 }

func getEventsById(contex *gin.Context){
	eventId,err:= strconv.ParseInt(contex.Param("id"),10,64)
	if err != nil {
		 contex.JSON(http.StatusBadRequest,gin.H{"message":"wrong eventId in params"})
		 return
	}
	event,err:=models.GetEventById(eventId)
    if err != nil {
		fmt.Print(err)
		contex.JSON(http.StatusBadRequest,gin.H{"message":"not able to fetch the eventbyid"})
		return;
	}
	contex.JSON(http.StatusOK,gin.H{"message":"event has been fetched by id","event":event})
}

func updateEvent(contex *gin.Context) {
	eventId,err:= strconv.ParseInt(contex.Param("id"),10,64)
	if err != nil {
		 contex.JSON(http.StatusBadRequest,gin.H{"message":"wrong eventId in params"})
		 return
	}

	_,err= models.GetEventById(eventId)
	
	if err != nil {
		contex.JSON(http.StatusBadRequest,gin.H{"message":"something went wrong in udate event"})
		return 
	}
	var updatedEvent models.Event
	updatedEvent.ID = eventId
	err = contex.ShouldBindJSON(&updatedEvent)
	if err!=nil {
		fmt.Print(err)
		contex.JSON(http.StatusBadRequest,gin.H{"message":" error in binding json in update event"})
		return 
	}
	err = updatedEvent.Update()
    if err != nil {
		fmt.Print(err)
		contex.JSON(http.StatusBadRequest,gin.H{"message":"failed to update event"})
		return 
	}
}

