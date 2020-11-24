package fuse

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Sheerley/pluggabl/codes"
	"github.com/Sheerley/pluggabl/convo"
	"github.com/Sheerley/pluggabl/db"
	"github.com/Sheerley/pluggabl/plog"
)

// Watchdog is used in servers main function to handle interrupt signal
func Watchdog(wg *sync.WaitGroup) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	plog.Debugf("watchdog is running")

	for {
		select {
		case <-c:
			plog.Warningf("recieved notification about interrupt signal")
			wg.Done()
			return
		default:
		}
		time.Sleep(time.Second)
	}
}

// Heartbeat is function that updates status of node in database every minute
func Heartbeat(conf convo.Config) {
	var err error
	count := 1
	for {
		err = db.UpdateTimestamp(context.Background(), conf)
		if err != nil {
			if count <= 5 {
				plog.Errorf("error #%v while updating timestamp: %v", count, err)
				count++
				time.Sleep(15 * time.Millisecond)
			} else {
				plog.Fatalf(codes.WorkerConnectionError,
					"fatal error while updating timestamp: %v", err)
			}
		} else {
			plog.Debugf("heartbeat updated")

			time.Sleep(50 * time.Second)
			count = 1
		}
	}
}
