package main

import (
	"flag"
	"fmt"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var argv arrayFlags

func main() {
	flag.Var(&argv, "arg", "Some description for this param.")
	flag.Parse()

	fmt.Printf("%+v\n", argv[3])
}
