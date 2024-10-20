package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e Event) Save() error {
	query := "INSERT INTO events (name, description, location, date_time, user_id) VALUES(?, ?, ?, ?, ?)"
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	exec, err := prepare.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id

	return nil
}

func (e Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?`
	prepare, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e Event) isAlreadyRegistered(userId int64) bool {
	query := "SELECT id FROM registrations WHERE event_id = ? AND user_id = ? LIMIT 1"
	row := db.DB.QueryRow(query, e.ID, userId)
	var event Event
	err := row.Scan(&event.ID)
	if err != nil {
		return false
	}

	return event.ID != 0
}

func (e Event) Register(userId int64) error {
	if e.isAlreadyRegistered(userId) {
		return nil
	}

	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return nil
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return nil
}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]Event, 0)

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
