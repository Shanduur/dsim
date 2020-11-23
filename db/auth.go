package db

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Sheerley/pluggabl/codes"
	"github.com/Sheerley/pluggabl/pb"
	"github.com/Sheerley/pluggabl/plog"
)

// Auth func is used to check if user is user credentials are correct
func Auth(ctx context.Context, user *pb.Credentials) error {
	var key []byte

	plog.Debugf("name: %v", user.GetUserId())

	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	err = conn.QueryRow(context.Background(), "SELECT user_key FROM users WHERE user_name = $1", user.UserId).Scan(&key)
	if err != nil {
		return fmt.Errorf("unable to execute querry: %v", err)
	}

	res := bytes.Compare(key, user.UserKey)

	if res != 0 {
		return &codes.NotAuthenticated{}
	}

	return nil
}
