package models

import (
	"gteruithi.com/demo-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	UserId      int64  `json:"user_id"`
}

func (e *Event) Save() error {
	queryInsert := `INSERT INTO events(name, description, location,  user_id) VALUES (?, ?, ?, ?) RETURNING id;`

	stmt, err := db.DB.Prepare(queryInsert)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
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
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserId)

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
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (event Event) Update() error {
	query := `UPDATE events SET name = ? , description = ?, location = ? WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID, event.Name, event.Description, event.Location)

	if err != nil {
		return err
	}

	return nil

}

func (event Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)

	if err != nil {
		return err
	}

	return nil

}
