package main

import (
	"context"
	"crypto/sha256"
	"time"

	"github.com/Sheerley/dsip/plog"
	"github.com/jackc/pgx/v4"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgresql://100.86.110.14:5432/dsipe_db?user=admin&password=XRB9UWu^bm8E^2aV")
	if err != nil {
		plog.Fatalf(1, "Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	var blobData []byte
	var blobID int

	blobData = []byte("Hello World")
	s256 := sha256.New()

	checksum := s256.Sum(blobData)

	dt := time.Now()

	parents := make([]int, 2)

	parents[0] = 1
	parents[1] = 2

	err = conn.QueryRow(context.Background(),
		"INSERT INTO blobs(blob_data, blob_type, blob_name, blob_checksum, insertion_date, parents)"+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING blob_id",
		blobData, 1, "name", checksum, dt.Format("2006-01-02"), parents).Scan(&blobID)
	if err != nil {
		plog.Fatalf(1, "QueryRow failed: %v", err)
	}

	plog.Messagef("OK")

	// permissions := 0666
	// var extension string
	// if blobType == 1 {
	// 	extension = ".txt"
	// } else {
	// 	extension = ""
	// }

	// err = ioutil.WriteFile("fetched/example"+extension, blobData, os.FileMode(permissions))
	// if err != nil {
	// 	plog.Fatalf(1, "Something went wrong while saving file: %v", err)
	// }
}
