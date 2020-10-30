package fuse

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Sheerley/pluggabl/pkg/plog"
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
