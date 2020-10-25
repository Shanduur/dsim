package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Sheerley/pluggabl/internal/convo"
)

func UpdateJobStatus(cfg convo.Config, value int) (err error) {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return
	}

	// in case of returning error rollback unfinished transaction
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "LOCK TABLE nodes IN ACCESS EXCLUSIVE MODE")
	if err != nil {
		return fmt.Errorf("unable to lock table: %v", err)
	}

	var count int

	err = tx.QueryRow(context.Background(), "SELECT COUNT(*) FROM nodes WHERE node_ip = $1").Scan(&count)
	if err != nil {
		return fmt.Errorf("unable to count in table: %v", err)
	}

	dt := time.Now()
	if count > 0 {
		err = tx.QueryRow(context.Background(), "SELECT node_reg_jobs FROM nodes WHERE node_ip = $1",
			cfg.WorkerAddress).Scan(&count)
		if err != nil {
			return fmt.Errorf("unable to get node_reg_jobs: %v", err)
		}

		err = tx.QueryRow(context.Background(), "UPDATE nodes SET node_reg_jobs = $1, node_timeout = $2 WHERE node_ip = $3",
			count+value, dt.Format("2006-01-02 15:04:05.070"), cfg.WorkerAddress).Scan(&count)
		if err != nil {
			return fmt.Errorf("unable to count in table: %v", err)
		}
	} else {
		return fmt.Errorf("node not registred")
	}

	tx.Commit(context.Background())

	return
}
