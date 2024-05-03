package routes

import (
	"net/http"
	"strconv"

	"example.com/event-planner/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events. Try again later."})

		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id."})
		return
	}
	event, err := models.GetEventById(int(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event, please try again."})
		return
	}

	context.JSON(http.StatusOK, event)
}
