package main

import (
	"github.com/JuKu/event-navigator-backend/db"
	"github.com/JuKu/event-navigator-backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db.InitDB()

	// for production usage (alternative: export GIN_MODE=release)
	// gin.SetMode(gin.ReleaseMode)

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080")

	if err != nil {
		return
	}
}

func getEvents(context *gin.Context) {
	events := model.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse Event model to create", "error_details": err.Error()})
	}

	// TODO: remove this later
	event.ID = 1
	event.CreatorID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
