package db

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/Sheerley/pluggabl/pkg/plog"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/pkg/pb"
)

// UploadFiles function inserts blobs into database
func UploadFiles(ctx context.Context, data [][]byte, fileInfo []*pb.FileInfo, user *pb.Credentials) (id []int, err error) {
	conn, err := connect()
	if err != nil {
		return append(id, codes.UnknownID), fmt.Errorf("unable to connect to database: %v", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return append(id, codes.UnknownID), err
	}

	// in case of returning error rollback unfinished transaction
	defer tx.Rollback(context.Background())

	count := 0
	dt := time.Now()
	ownerID := 0

	err = conn.QueryRow(context.Background(), "SELECT user_id FROM users WHERE user_name = $1", user.UserId).Scan(&ownerID)
	if err != nil {
		return append(id, codes.UnknownID), fmt.Errorf("unable to execute querry: %v", err)
	}

	for i := 0; i < len(data); i++ {
		var blobID int

		err = conn.QueryRow(context.Background(), "SELECT COUNT(type_id) FROM filetypes WHERE type_extension = $1",
			fileInfo[i].FileExtension).Scan(&count)
		if err != nil {
			return append(id, codes.UnknownID), fmt.Errorf("unable to execute querry: %v", err)
		}

		typeID := 1
		if count >= 1 {
			err = conn.QueryRow(context.Background(), "SELECT type_id FROM filetypes WHERE type_extension = $1 LIMIT 1",
				fileInfo[i].FileExtension).Scan(&typeID)
			if err != nil {
				return append(id, codes.UnknownID), fmt.Errorf("unable to execute querry: %v", err)
			}
		}

		err = tx.QueryRow(context.Background(),
			"INSERT INTO blobs(blob_data, blob_type, blob_name, owner_id, insertion_date)"+
				"VALUES ($1, $2, $3, $4, $5) RETURNING blob_id",
			data[i], typeID, fmt.Sprint(i), ownerID, dt.Format("2006-01-02")).Scan(&blobID)

		plog.Debugf("id returned: %v", blobID)

		if err != nil {
			return append(id, codes.UnknownID), err
		}

		id = append(id, blobID)
	}

	if ctx.Err() == context.Canceled {
		return append(id, codes.UnknownID), &codes.SignalCanceled{}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		id = nil
		return append(id, codes.UnknownID), err
	}

	return id, nil
}

// UserExists checks if user exists inside database
func UserExists(user *pb.Credentials) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	count := 0

	err = conn.QueryRow(context.Background(), "SELECT COUNT(user_id) FROM users WHERE user_name = $1", user.UserId).Scan(&count)
	if err != nil {
		return fmt.Errorf("unable to execute querry: %v", err)
	}

	if count != 0 {
		return &codes.RecordExists{}
	}

	return nil
}

// GetFile retrieves result blob from database
func GetFile(ctx context.Context, id int64) (result []byte, extension string, err error) {
	conn, err := connect()
	if err != nil {
		err = fmt.Errorf("unable to connect to database: %v", err)
		return
	}

	err = conn.QueryRow(context.Background(), "SELECT blob_data FROM blobs WHERE blob_id = $1", id).Scan(&result)
	if err != nil {
		err = fmt.Errorf("unable to execute querry: %v", err)
		return
	}

	var typeID int64

	err = conn.QueryRow(context.Background(), "SELECT blob_type FROM blobs WHERE blob_id = $1", id).Scan(&typeID)
	if err != nil {
		err = fmt.Errorf("unable to execute querry: %v", err)
		return
	}

	err = conn.QueryRow(context.Background(), "SELECT type_extension FROM filetypes WHERE type_id = $1", typeID).Scan(&extension)
	if err != nil {
		err = fmt.Errorf("unable to execute querry: %v", err)
		return
	}

	return
}

// CreateUser inserts new user into table
func CreateUser(ctx context.Context, user *pb.Credentials) error {
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

	if ctx.Err() == context.Canceled {
		return &codes.SignalCanceled{}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes user from table
func DeleteUser(ctx context.Context, user *pb.Credentials) error {
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

	if ctx.Err() == context.Canceled {
		return &codes.SignalCanceled{}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
