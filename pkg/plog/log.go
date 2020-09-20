package plog

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sheerley/pluggabl/internal/codes"
)

const dtFormat = "01-02-2006 15:04:05.000"

// Constants describing level of log:
// - INFO - all logs will be registred,
// - WARNING - only warning logs and above will be registred,
// - ERROR - only error logs and above will be registred,
// - FATAL - (NOT ADVISED!) only fatal logs will be registred.
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

	return fmt.Sprintf("%v [%v] ", dt.Format(dtFormat), s)
}

// CloseLogFile is used to defer closing the log file.
func CloseLogFile() {
	err := outFile.Close()
	if err != nil {
		Errorf("%v", err)
	}
}

// SetLogLevel is used to set the amount of infromation that should be logged.
func SetLogLevel(level int) {
	logLevel = level
}

// SetLogFile is used to specify the location of log file.
func SetLogFile(path string) {
	var err error
	outFile, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		outFile = os.Stderr
		Errorf("%v\n", err)
	}
}

// Messagef is used to create formatted message - INFO level.
func Messagef(format string, v ...interface{}) {
	if logLevel <= INFO {
		_, err := fmt.Fprintf(outFile, getHeader(INFO)+format, v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Errorf is used to create formatted message - ERROR level.
func Errorf(format string, v ...interface{}) {
	if logLevel <= ERROR {
		_, err := fmt.Fprintf(outFile, getHeader(ERROR)+format, v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Warningf is used to create formatted message - WARNING level.
func Warningf(format string, v ...interface{}) {
	if logLevel <= WARNING {
		_, err := fmt.Fprintf(outFile, getHeader(WARNING)+format, v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Fatalf is used to create formatted exit message - FATAL level.
// Warning! It forces exit of the app with exit code provided as first function argument
func Fatalf(code int, format string, v ...interface{}) {
	_, err := fmt.Fprintf(outFile, getHeader(FATAL)+format, v...)
	if err != nil {
		log.Printf("%v", err)
		os.Exit(codes.LogError)
	}
	os.Exit(code)
}
