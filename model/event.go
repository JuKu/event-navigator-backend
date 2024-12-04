package model

import (
	"fmt"
	"github.com/JuKu/event-navigator-backend/db"
	"time"
)

// Event is a struct which is returned to the client by the server
type Event struct {
	ID          int64  `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Organizer   string `json:"organizer" binding:"required"`
	// date and time
	DateTime     time.Time `json:"datetime" binding:"required"`
	CalendarWeek int       `json:"calendar_week"`
	Year         int       `json:"year"`
	// creator
	CreatorID int64 `json:"creator_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (
                    title, description, location, organizer, datetime, calendar_week, year, creator_id
		) VALUES (
		          ?, ?, ?, ?, ?, ?, ?, ?
		)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error while preparing statement", err)
		return err
	}

	defer stmt.Close()

	// calculate calcandar week
	year, week := calculateGetCalendarWeekAndYear(e.DateTime)
	e.Year = year
	e.CalendarWeek = week

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.Organizer, e.DateTime, e.CalendarWeek, e.Year, e.CreatorID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Organizer, &event.DateTime, &event.CalendarWeek, &event.Year, &event.CreatorID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(eventID int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id=?`
	row := db.DB.QueryRow(query, eventID)

	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Organizer, &event.DateTime, &event.CalendarWeek, &event.Year, &event.CreatorID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `UPDATE events
	SET title = ?, description = ?, location = ?, organizer = ?, datetime = ?, calendar_week = ?, year = ?, creator_id = ?
    WHERE id=?`

	year, week := calculateGetCalendarWeekAndYear(event.DateTime)

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.Organizer, event.DateTime, week, year, event.CreatorID, event.ID)

	return err
}

func calculateGetCalendarWeekAndYear(timestamp time.Time) (year, week int) {
	return timestamp.ISOWeek()
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)

	return err
}
