package main

import (
	"github.com/Sheerley/pluggabl/pkg/plog"
	"time"
)

func main() {
	plog.SetLogFile("cmd/plog-example/log.txt")
	defer plog.CloseLogFile() // will not run due to Fatalf

	plog.Messagef("%v %v %v\n", 12, 13, 14)
	time.Sleep(2 * time.Second)

	plog.Warningf("%v %v %v\n", 12, 13, 14)
	time.Sleep(32 * time.Millisecond)

	plog.Errorf("%v %v %v\n", 12, 13, 14)

	plog.SetLogLevel(plog.ERROR)

	plog.Messagef("%v %v %v\n", 12, 13, 14)
	plog.Warningf("%v %v %v\n", 12, 13, 14)
	time.Sleep(12 * time.Millisecond)

	plog.Fatalf(1,"%v %v %v\n", 12, 13, 14)
}
