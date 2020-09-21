package main

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/Sheerley/pluggabl/pkg/plog"
	"github.com/jackc/pgx/v4"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgresql://localhost/database?user=user&password=password")
	if err != nil {
		plog.Fatalf(1, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var blobData []byte
	var blobType int32

	err = conn.QueryRow(context.Background(), "SELECT blob_data, blob_type FROM blobs WHERE blob_id=$1", 1).Scan(&blobData, &blobType)
	if err != nil {

		
		
		plog.Fatalf(1, "QueryRow failed: %v\n", err)
	}

	permissions := 0666
	var extension string
	if blobType == 1 {
		extension = ".txt"
	} else {
		extension = ""
	}

	err = ioutil.WriteFile("fetched/example"+extension, blobData, os.FileMode(permissions))
	if err != nil {
		plog.Fatalf(1, "Something went wrong while saving file: %v\n", err)
	}
}
