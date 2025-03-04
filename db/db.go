package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		fmt.Println(err)
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {

	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		fmt.Println(err)
		panic("could not create user table")
	}

	createEventsTable :=
		`
			CREATE TABLE IF NOT EXISTS events (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					name TEXT NOT NULL,
					description TEXT NOT NULL,
					location TEXT NOT NULL,
					user_id INTEGER,
					FOREIGN KEY(user_id) REFERENCES users (id)
			);
		`

	_, err = DB.Query(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("could not create event table")
	}
}
