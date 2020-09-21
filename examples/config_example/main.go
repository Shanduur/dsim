package main

import (
	"fmt"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/internal/convo"
	"github.com/Sheerley/pluggabl/pkg/plog"
)

func main() {
	location := "default"
	err := convo.LoadConfiguration(location)
	if err != nil {
		plog.Fatalf(codes.ConfError, "Failed to load configuration from %v: %v", location, err)
	}

	fmt.Println(convo.Configuration)
}
