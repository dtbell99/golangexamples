package localdatabase

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // Using loader
)

// LogMessage : Represents a log entry
type LogMessage struct {
	Message string `json:"message"`
}

// LogEntry : Represents database log entries
type LogEntry struct {
	ID      int       `json:"id"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// FindAllLogMessage : Function returning all logs in database
func FindAllLogMessage() []LogEntry {
	logs := make([]LogEntry, 0)
	db, err := sql.Open("sqlite3", "./localdatabase.db")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM log")
	defer rows.Close()

	checkErr(err)
	var id int
	var msg string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&id, &msg, &created)
		checkErr(err)
		nl := LogEntry{ID: id, Message: msg, Created: created}
		logs = append(logs, nl)
	}

	return logs
}

// AddLogMessage : Allows records to be added to the sqlite log table
func AddLogMessage(message string) int64 {

	db, err := sql.Open("sqlite3", "./localdatabase.db")
	checkErr(err)

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO log (MESSAGE) values(?)")
	checkErr(err)

	res, err := stmt.Exec(message)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Printf("id created: %d", id)
	return id
}
