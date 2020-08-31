package plog

import (
	"os"
	"strings"
	"testing"
)

func TestSetLogLevel(t *testing.T) {
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
}

func TestGetHeader(t *testing.T) {
	vals := make([]string, 4)

	vals[0] = "INFO"
	vals[1] = "WARN"
	vals[2] = "ERROR"
	vals[3] = "FATAL"

	var i int8

	for i = 0; i <= 3; i++ {
		h := getHeader(i)
		out := strings.Contains(h, vals[i])
		if out != true {
			t.Errorf("getHeader: got %v, want %v", vals[i], h)
		}
	}
}
