package plog

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestSSetLogLevel(t *testing.T) {
	levels := []int{VERBOSE, DEBUG, INFO, WARNING, ERROR, FATAL}
	strs := []string{"VERBOSE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
	data := make(map[int]string)

	for i, l := range levels {
		data[l] = strs[i]
	}

	for _, l := range levels {
		SSetLogLevel(data[l])

		if logLevel != l {
			t.Errorf("SSetLogLevel: got %d, want %d", logLevel, l)
		}
	}
}

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(DEBUG)
	if logLevel != DEBUG {
		t.Errorf("SetLogLevel: got %d, want %d", logLevel, ERROR)
	}

	SetLogLevel(VERBOSE)
	if logLevel != VERBOSE {
		t.Errorf("SetLogLevel: got %d, want %d", logLevel, ERROR)
	}

	SetLogLevel(ERROR)
	if logLevel != ERROR {
		t.Errorf("SetLogLevel: got %d, want %d", logLevel, ERROR)
	}

	SetLogLevel(WARNING)
	if logLevel != WARNING {
		t.Errorf("SetLogLevel: got %d, want %d", logLevel, WARNING)
	}

	SetLogLevel(FATAL)
	if logLevel != FATAL {
		t.Errorf("SetLogLevel: got %d, want %d", logLevel, FATAL)
	}

	SetLogLevel(INFO)
	if logLevel != INFO {
		t.Errorf("SetLogLevel: got %d, want %d", logLevel, INFO)
	}
}

func TestLogFile(t *testing.T) {
	path := "./log.tmp.txt"

	f, err := os.Create(path)
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}
	f.Close()

	SetLogFile(path)
	if outFile.Name() != path {
		t.Errorf("SetLogLevel: got %v, want %v", outFile.Name(), path)
	}

	CloseLogFile()
	CloseLogFile()
}

func TestGetHeader(t *testing.T) {
	vals := make([]string, 6)

	vals[0] = "VERBOSE"
	vals[1] = "DEBUG"
	vals[2] = "INFO"
	vals[3] = "WARN"
	vals[4] = "ERROR"
	vals[5] = "FATAL"

	var i int8

	for i = 0; i <= 3; i++ {
		h := getHeader(i)
		out := strings.Contains(h, vals[i])
		if out != true {
			t.Errorf("getHeader: got %v, want %v", vals[i], h)
		}
	}
}

func TestFunctions(t *testing.T) {
	tMessage := "test message"

	SetLogLevel(VERBOSE)

	Messagef("%v", tMessage)
	Warningf("%v", tMessage)
	Errorf("%v", tMessage)
	Debugf("%v", tMessage)
	Verbosef("%v", tMessage)
	Verbose(tMessage)

	ctx, cancel := context.WithCancel(context.Background())

	defer ContextStatus(ctx)

	go func(cf context.CancelFunc) {
		cf()
	}(cancel)
}

func TestSplash(t *testing.T) {
	Splash("TEST_STRING")
}
