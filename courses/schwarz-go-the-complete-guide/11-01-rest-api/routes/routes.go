package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticatedRoutesGroup := server.Group("/")
	authenticatedRoutesGroup.Use(middlewares.Authenticate)
	authenticatedRoutesGroup.POST("/events", createEvent)
	authenticatedRoutesGroup.PUT("/events/:id", updateEvent)
	authenticatedRoutesGroup.DELETE("/events/:id", deleteEvent)
	authenticatedRoutesGroup.POST("/events/:id/register", registerToEvent)
	authenticatedRoutesGroup.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
