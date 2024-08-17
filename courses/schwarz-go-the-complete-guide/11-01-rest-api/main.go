package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/db"
	"rest-api/models"
	"strconv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", func(c *gin.Context) {
		events, err := models.GetEvents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("DB error fetching events: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"events": events,
		})
	})
	server.GET("/events/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("ID param error: %s", err.Error()),
			})
			return
		}
		event, err := models.GetEventById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("DB error fetching event: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"event": event,
		})
	})
	server.POST("/events", func(c *gin.Context) {
		var event models.Event
		err := c.ShouldBindJSON(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Marshalling error: %s", err.Error()),
			})
			return
		}

		err = event.Save()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error while saving event: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while saving event"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Event created", "event": event,
		})
	})

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
