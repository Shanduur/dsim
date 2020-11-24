package main

import (
	"flag"

	"github.com/Sheerley/pluggabl/codes"
	"github.com/Sheerley/pluggabl/compute"
	"github.com/Sheerley/pluggabl/plog"
)

func main() {
	plog.SetLogLevel(plog.INFO)
	query := flag.String("query", "", "name of query file")
	train := flag.String("train", "", "name of train file")
	out := flag.String("out", "", "name of output file")
	logLevel := flag.Int("log-level", plog.INFO, "log level")

	flag.Parse()
	if len(*query) == 0 || len(*train) == 0 || len(*out) == 0 {
		plog.Fatalf(codes.IncorrectArgs, "error with parsing arguments")
	}

	plog.SetLogLevel(*logLevel)

	plog.Messagef("running with args:\n -query=%v,\n -train=%v,\n -out=%v,\n -log-level=%v", *query, *train, *out, *logLevel)

	outMat := compute.BestMatches(*query, *train)

	plog.Debugf("saving output [empty=%v]", outMat.Empty())
	compute.SaveMat(*out, outMat)
}
