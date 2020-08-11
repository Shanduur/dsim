package main

import (
	"fmt"
	"github.com/Sheerley/pluggabl/internal/convo"
	"log"
)

func main() {
	location := "default"
	err := convo.LoadConfiguration(location)
	if err != nil {
		log.Fatalf("Failed to load configuration from %v: %v", location ,err)
	}

	fmt.Println(convo.Configuration)
}
