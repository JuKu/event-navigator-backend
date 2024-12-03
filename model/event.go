package model

import "time"

// Event is a struct which is returned to the client by the server
type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Organizer   string `json:"organizer" binding:"required"`
	// date and time
	DateTime     time.Time `json:"datetime" binding:"required"`
	CalendarWeek int       `json:"calendar_week"`
	Year         int       `json:"year"`
	// creator
	CreatorID int `json:"creator_id"`
}

var events = []Event{}

func (e Event) Save() {
	// TODO: add event to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
