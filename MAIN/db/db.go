package db

import (
	// run Init function in package mysql
	"database/sql"
	"log"
	"os"

	// run Init function in package mysql
	_ "github.com/go-sql-driver/mysql"
)

// Connection is a database handle representing a pool of zero or more
type Connection struct {
	Server *sql.DB
}

// Connect opens a database specified by its database driver name and a
// driver-specific data source name, usually consisting of at least a
// database name and connection information.
func Connect() Connection {
	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE"))
	if err != nil {
		log.Print(err.Error())
	}
	return Connection{db}
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server
// to finish.
func (c Connection) Close() {
	c.Server.Close()
}

// Read single row
func (c Connection) Read(stmnt string) *sql.Row {
	s := c.Server
	return s.QueryRow(stmnt)
}

// Query multiple rows
func (c Connection) Query(stmnt string) (*sql.Rows, error) {
	s := c.Server
	return s.Query(stmnt)
}
