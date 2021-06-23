package plog

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Sheerley/dsip/codes"
)

const dtFormat = "01-02-2006 15:04:05.000"

// Constants describing level of log:
// - DEBUG - additional debug logs will be registred,
// - INFO - all logs will be registred,
// - WARNING - only warning logs and above will be registred,
// - ERROR - only error logs and above will be registred,
// - FATAL - (NOT ADVISED!) only fatal logs will be registred.
const (
	VERBOSE = iota
	DEBUG   = iota
	INFO    = iota
	WARNING = iota
	ERROR   = iota
	FATAL   = iota
)

var outFile = os.Stderr
var logLevel = INFO

func getHeader(level int8) string {
	dt := time.Now()

	var s string

	switch level {
	case VERBOSE:
		s = "VERBOSE"
	case DEBUG:
		s = "DEBUG"
	case INFO:
		s = "INFO"
	case WARNING:
		s = "WARNING"
	case ERROR:
		s = "ERROR"
	case FATAL:
		s = "FATAL"
	}

	return fmt.Sprintf("%v [%v] ", dt.Format(dtFormat), s)
}

// CloseLogFile is used to defer closing the log file.
func CloseLogFile() {
	var err error

	if outFile != os.Stderr {
		err = outFile.Close()
		outFile = os.Stderr
	}

	if err != nil {
		Errorf("%v", err)
	}
}

// SetLogLevel is used to set the amount of infromation that should be logged.
func SetLogLevel(level int) {
	logLevel = level
}

// GetLogLevel returns human readable level description
func GetLogLevel() string {
	switch logLevel {
	case VERBOSE:
		return "VERBOSE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return ""
	}
}

// SSetLogLevel sets the level of logs according to the value from the string
func SSetLogLevel(level string) {
	level = strings.ToUpper(level)

	switch level {
	case "VERBOSE":
		SetLogLevel(VERBOSE)
	case "DEBUG":
		SetLogLevel(DEBUG)
	case "INFO":
		SetLogLevel(INFO)
	case "WARNING":
		SetLogLevel(WARNING)
	case "ERROR":
		SetLogLevel(ERROR)
	case "FATAL":
		SetLogLevel(FATAL)
	}
}

// SetLogFile is used to specify the location of log file.
func SetLogFile(path string) {
	var err error
	outFile, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		outFile = os.Stderr
		Errorf("%v", err)
	}
}

// Messagef is used to create formatted message - INFO level.
func Messagef(format string, v ...interface{}) {
	if logLevel <= INFO {
		_, err := fmt.Fprintf(outFile, getHeader(INFO)+format+"\n", v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Errorf is used to create formatted message - ERROR level.
func Errorf(format string, v ...interface{}) {
	if logLevel <= ERROR {
		_, err := fmt.Fprintf(outFile, getHeader(ERROR)+format+"\n", v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Warningf is used to create formatted message - WARNING level.
func Warningf(format string, v ...interface{}) {
	if logLevel <= WARNING {
		_, err := fmt.Fprintf(outFile, getHeader(WARNING)+format+"\n", v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Debugf is used to create formatted debug message - DEBUG level.
func Debugf(format string, v ...interface{}) {
	if logLevel <= DEBUG {
		_, err := fmt.Fprintf(outFile, getHeader(DEBUG)+format+"\n", v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Verbosef is used to create extended formatted debug message - VERBOSE level.
func Verbosef(format string, v ...interface{}) {
	if logLevel <= VERBOSE {
		_, err := fmt.Fprintf(outFile, getHeader(VERBOSE)+format+"\n", v...)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(codes.LogError)
		}
	}
}

// Verbose displays all info about variable
func Verbose(v interface{}) {
	Verbosef("requested variable info:\n"+
		"\t- %+v\n"+
		"\t- %T",
		v, v)
}

// Fatalf is used to create formatted exit message - FATAL level.
// Warning! It forces exit of the app with exit code provided as first function argument
func Fatalf(code int, format string, v ...interface{}) {
	_, err := fmt.Fprintf(outFile, getHeader(FATAL)+format+"\n", v...)
	if err != nil {
		log.Printf("%v", err)
		os.Exit(codes.LogError)
	}
	os.Exit(code)
}

// ContextStatus creates warning in logs when ctx recievs cancelation.
func ContextStatus(ctx context.Context) {
	if ctx.Err() == context.Canceled {
		Warningf("%v", codes.ErrSignalCanceled)
	}
}

// Splash is used to display informations on boot
func Splash(s string) {
	splash := "" +
		" ________  ___       ___  ___  ________  ________  ________  ________  ___           \n" +
		"|\\   __  \\|\\  \\     |\\  \\|\\  \\|\\   ____\\|\\   ____\\|\\   __  \\|\\   __  \\|\\  \\          \n" +
		"\\ \\  \\|\\  \\ \\  \\    \\ \\  \\\\\\  \\ \\  \\___|\\ \\  \\___|\\ \\  \\|\\  \\ \\  \\|\\ /\\ \\  \\         \n" +
		" \\ \\   ____\\ \\  \\    \\ \\  \\\\\\  \\ \\  \\  __\\ \\  \\  __\\ \\   __  \\ \\   __  \\ \\  \\        \n" +
		"  \\ \\  \\___|\\ \\  \\____\\ \\  \\\\\\  \\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\ \\  \\ \\  \\|\\  \\ \\  \\____   \n" +
		"   \\ \\__\\    \\ \\_______\\ \\_______\\ \\_______\\ \\_______\\ \\__\\ \\__\\ \\_______\\ \\_______\\ \n" +
		"    \\|__|     \\|_______|\\|_______|\\|_______|\\|_______|\\|__|\\|__|\\|_______|\\|_______| \n" +
		"                                                                                     \n"
	fmt.Fprintf(outFile, "%v\n%v\n\t%v %v\n", splash, s,
		"Your Log Level is set to:", GetLogLevel())
}
