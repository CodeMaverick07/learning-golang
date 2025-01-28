package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/events",getAllEvents)
	server.POST("/events",createEvents)
	server.GET("/events/:id",getEventsById)
	server.PUT("/events/:id",updateEvent) 
	server.DELETE("/events/:id",deleteEvent)
}