package models

import (
	"context"
	"time"

	"example.com/event-planner/db"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, datetime, user_id)
	VALUES($1,$2,$3,$4,$5)
	RETURNING id
`
	id := 0
	err := db.DB.QueryRow(context.Background(), query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&id)

	if err != nil {
		return err
	}

	e.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var lEvents []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		lEvents = append(lEvents, event)
	}

	return lEvents, nil
}

func GetEventById(id int) (*Event, error) {
	query := `SELECT * FROM events WHERE id=$1`
	var event Event

	row := db.DB.QueryRow(context.Background(), query, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, dateTime = $4
		WHERE id = $5
	`

	_, err := db.DB.Exec(context.Background(), query, event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := db.DB.Exec(context.Background(), query, event.ID)

	return err
}
