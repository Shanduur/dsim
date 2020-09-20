package db

import "github.com/jackc/pgx/v4"

// Conn is default connection to database
var Conn pgx.Conn

// Connect is used to connect to database
func Connect() error {

	return nil
}

// Close is meant to be deffered and close existing database connection
func Close() {

}
