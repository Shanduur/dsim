package db

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Sheerley/pluggabl/pkg/pb"
)

// UserExists checks if user exists inside database
func UserExists(user *pb.Credentials) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	count := 0

	err = conn.QueryRow(context.Background(), "SELECT COUNT(user_id) FROM users WHERE user_name =  $1", user.UserId).Scan(&count)
	if err != nil {
		return fmt.Errorf("unable to execute querry: %v", err)
	}

	if count != 0 {
		return fmt.Errorf("querry returned more than 0")
	}

	return nil
}

// CreateUser inserts new user into table
func CreateUser(user *pb.Credentials) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	// in case of returning error rollback unfinished transaction
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO users(user_name, user_key) VALUES ($1, $2)", user.UserId, user.UserKey)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes user from table
func DeleteUser(user *pb.Credentials) error {
	var key []byte

	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	err = conn.QueryRow(context.Background(), "SELECT user_key FROM users WHERE user_name = $1", user.UserId).Scan(&key)
	if err != nil {
		return fmt.Errorf("unable to execute querry: %v", err)
	}

	res := bytes.Compare(key, user.UserKey)

	if res != 0 {
		return fmt.Errorf("passwords are not equal")
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	// in case of returning error rollback unfinished transaction
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO users(user_name, user_key) VALUES ($1, $2)", user.UserId, user.UserKey)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
