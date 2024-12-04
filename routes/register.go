package routes

import (
	"github.com/JuKu/event-navigator-backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error": err.Error()})
	}

	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event by id.", "error": err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registered for event.", "event": event})
}

func cancelRegistrationForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error": err.Error()})
	}

	var event model.Event
	event.ID = eventId

	/*event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event by id.", "error": err.Error()})
		return
	}*/

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User cancelled registration for event.", "event": event})
}
