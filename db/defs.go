package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Sheerley/dsim/codes"
	"github.com/Sheerley/dsim/convo"
	"github.com/Sheerley/dsim/plog"

	"github.com/jackc/pgx/v4"
)

var dbURL = "postgresql://%v:%v/%v?user=%v&password=%v"

type dbconf struct {
}

func connect() (*pgx.Conn, error) {
	url := os.Getenv("PG_DATABASE")

	if len(url) < len(dbURL) {
		plog.Messagef("env var PG_DATABASE is not valid, attempting to load config file")

		conf, err := convo.LoadConfiguration("/etc/dsim/config_db.json")
		if err != nil {
			plog.Fatalf(codes.ConfError, "error while loading db configuration: %v", err)
		} else {
			url = fmt.Sprintf(dbURL, conf.Address, conf.Port, conf.DatabaseName,
				conf.DatabaseUsername, conf.DatabasePassword)
		}
	}

	return pgx.Connect(context.Background(), url)
}
