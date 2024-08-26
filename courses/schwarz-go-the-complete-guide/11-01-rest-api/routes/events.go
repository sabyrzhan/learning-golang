package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"strconv"
)

func getEventById(c *gin.Context) {
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
}

func getEvents(c *gin.Context) {
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
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Marshalling error: %s", err.Error()),
		})
		return
	}

	event.UserID = c.GetInt64("userId")

	err = event.Save()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while saving event: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while saving event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created", "event": event,
	})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("ID param error: %s", err.Error()),
		})
		return
	}
	_, err = models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("DB error fetching event: %s", err.Error()),
		})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Marshalling error: %s", err.Error()),
		})

		return
	}
	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Update event error: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully", "event": updatedEvent,
	})
}

func deleteEvent(c *gin.Context) {
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

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Delete event error: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
