package plog

import (
	"fmt"
	"log"
	"os"
	"time"
)

const dtFormat = "01-02-2006 15:04:05.000"

const (
	INFO    = iota
	WARNING = iota
	ERROR   = iota
	FATAL   = iota
)

var outFile = os.Stderr
var logLevel = INFO

func getHeader(lvl int8) string {
	dt := time.Now()

	var s string

	if lvl == 0 {
		s = "INFO"
	} else if lvl == 1 {
		s = "WARN"
	} else if lvl == 2 {
		s = "ERROR"
	} else {
		s = "FATAL"
	}

	return fmt.Sprintf("[%v] %v ", dt.Format(dtFormat), s)
}

func CloseLogFile() {
	err := outFile.Close()
	if err != nil {
		Errorf("%v", err)
	}
}

func SetLogLevel(level int) {
	logLevel = level
}

func SetLogFile(path string) {
	var err error
	outFile, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		outFile = os.Stderr
		Errorf("%v\n", err)
	}
}

func Messagef(format string, v ...interface{}) {
	if logLevel <= INFO {
		_, err := fmt.Fprintf(outFile, getHeader(INFO)+format, v...)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel <= ERROR {
		_, err := fmt.Fprintf(outFile, getHeader(ERROR)+format, v...)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func Warningf(format string, v ...interface{}) {
	if logLevel <= WARNING {
		_, err := fmt.Fprintf(outFile, getHeader(WARNING)+format, v...)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func Fatalf(code int, format string, v ...interface{}) {
	_, err := fmt.Fprintf(outFile, getHeader(FATAL)+format, v...)
	if err != nil {
		log.Fatalf("%v", err)
	}
	os.Exit(code)
}
