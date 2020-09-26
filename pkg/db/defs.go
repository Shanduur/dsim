package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

func connect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}

func close() {

}
