package main

import (
	"github.com/Sheerley/dsim/codes"
	"github.com/Sheerley/dsim/convo"
	"github.com/Sheerley/dsim/plog"
)

func main() {
	location := "config/config_secondary.json"
	conf, err := convo.LoadConfiguration(location)
	if err != nil {
		plog.Fatalf(codes.ConfError, "Failed to load configuration from %v: %v", location, err)
	}

	plog.Messagef("%v", conf)
}
