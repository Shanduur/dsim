package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Sheerley/pluggabl/internal/convo"
)

// RegisterNode is used in the start of server to register node inside database
func RegisterNode(cfg convo.Config) (err error) {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return
	}

	var count int

	err = tx.QueryRow(context.Background(), "SELECT COUNT(*) FROM nodes WHERE node_ip = $1",
		fmt.Sprintf("%v", cfg.SecondaryNodeAddress)).Scan(&count)
	if err != nil {
		return fmt.Errorf("unable to count in table: %v", err)
	}

	dt := time.Now()
	if count == 0 {
		_, err = tx.Exec(context.Background(),
			"INSERT INTO nodes(node_ip, node_port, node_reg_jobs, node_max_jobs, node_timeout) VALUES($1, $2, 0, $3, $4)",
			fmt.Sprintf("%v", cfg.SecondaryNodeAddress), cfg.SecondaryNodePort, 4, dt.Format("2006-01-02 15:04:05.070"))
		if err != nil {
			return fmt.Errorf("unable to insert node config into table: %v", err)
		}
	} else {
		return fmt.Errorf("node already registred")
	}

	tx.Commit(context.Background())

	return
}

// UnregisterNode is used in the end of server execution to remove node from database
func UnregisterNode(cfg convo.Config) (err error) {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return
	}

	var count int

	err = tx.QueryRow(context.Background(), "SELECT COUNT(*) FROM nodes WHERE node_ip = $1",
		fmt.Sprintf("%v", cfg.SecondaryNodeAddress)).Scan(&count)
	if err != nil {
		return fmt.Errorf("unable to count in table: %v", err)
	}

	if count > 0 {
		_, err = tx.Exec(context.Background(), "DELETE FROM nodes WHERE node_ip = $1", fmt.Sprintf("%v", cfg.SecondaryNodeAddress))
		if err != nil {
			return fmt.Errorf("unable to insert node config into table: %v", err)
		}
	} else {
		return fmt.Errorf("node already registred")
	}

	tx.Commit(context.Background())

	return
}

// UpdateJobStatus is used to update status of job inside node config table
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

	err = tx.QueryRow(context.Background(), "SELECT COUNT(*) FROM nodes WHERE node_ip = $1",
		fmt.Sprintf("%v", cfg.SecondaryNodeAddress)).Scan(&count)
	if err != nil {
		return fmt.Errorf("unable to count in table: %v", err)
	}

	dt := time.Now()
	if count > 0 {
		err = tx.QueryRow(context.Background(), "SELECT node_reg_jobs FROM nodes WHERE node_ip = $1",
			fmt.Sprintf("%v", cfg.SecondaryNodeAddress)).Scan(&count)
		if err != nil {
			return fmt.Errorf("unable to get node_reg_jobs: %v", err)
		}

		_, err = tx.Exec(context.Background(), "UPDATE nodes SET node_reg_jobs = $1, node_timeout = $2 WHERE node_ip = $3",
			count+value, dt.Format("2006-01-02 15:04:05.070"), fmt.Sprintf("%v", cfg.SecondaryNodeAddress))
		if err != nil {
			return fmt.Errorf("unable to update node in table: %v", err)
		}
	} else {
		return fmt.Errorf("node not registred")
	}

	tx.Commit(context.Background())

	return
}
