package routes

import (
	"example.com/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events",getAllEvents)
    Authenticated:= server.Group("/")
	Authenticated.Use(middlewares.Authenticate)
	Authenticated.POST("/events",createEvents)
	Authenticated.PUT("/events/:id",updateEvent)
	Authenticated.DELETE("/events/:id",deleteEvent)
	Authenticated.POST("/events/:id/register",registerEvent)
	Authenticated.DELETE( "/events/:id/unregister",unregisterEvent)
	server.GET("/events/:id",getEventsById)
	server.POST("/signup",signup)
	server.POST("/login",login)
}