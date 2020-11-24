package main

import (
	"github.com/Sheerley/pluggabl/codes"
	"github.com/Sheerley/pluggabl/convo"
	"github.com/Sheerley/pluggabl/plog"
)

func main() {
	location := "config/config_secondary.json"
	conf, err := convo.LoadConfiguration(location)
	if err != nil {
		plog.Fatalf(codes.ConfError, "Failed to load configuration from %v: %v", location, err)
	}

	plog.Messagef("%v", conf)
}
