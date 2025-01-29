package routes

import (
	"net/http"
	"strconv"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context ) {
	userId := context.GetInt64("userId")
  eventId,err := strconv.ParseInt(context.Param("id"),10,64)
  if err != nil {
	context.JSON(http.StatusBadRequest,gin.H{"message":"wrong eventId in params"})
	return
  }
  event,err := models.GetEventById(eventId)
  if err != nil {
	context.JSON(http.StatusBadRequest,gin.H{"message":"not able to fetch the eventbyid"})
	return
  }
  err= event.Register(userId)
  if err != nil {
	context.JSON(http.StatusBadRequest,gin.H{"message":"not able to register"})
	return
  } 
  context.JSON(http.StatusOK,gin.H{"message":"registered successfully"})

}

func unregisterEvent(context *gin.Context ) {
	userId := context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"),10,64)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"wrong eventId in params"})
		return
	}
	event,err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"not able to fetch the eventbyid"})
		return
	}
	err= event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"not able to unregister"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"unregistered successfully"})

}