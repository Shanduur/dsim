package main

import (
	"flag"

	"github.com/Sheerley/pluggabl/codes"
	"github.com/Sheerley/pluggabl/compute"
	"github.com/Sheerley/pluggabl/plog"
)

func main() {
	plog.SetLogLevel(plog.INFO)
	img1 := flag.String("img1", "", "name of first image file")
	img2 := flag.String("img2", "", "name of second image file")
	out := flag.String("out", "", "name of output file")
	logLevel := flag.Int("log-level", plog.INFO, "log level")

	flag.Parse()
	if len(*img1) == 0 || len(*img2) == 0 || len(*out) == 0 {
		plog.Fatalf(codes.IncorrectArgs, "error with parsing arguments")
	}

	plog.SetLogLevel(*logLevel)

	plog.Messagef("running with args:\n -img1=%v,\n -img2=%v,\n -out=%v,\n -log-level=%v", *img1, *img2, *out, *logLevel)

	outMat := compute.BestMatches(*img1, *img2)

	plog.Debugf("saving output [empty=%v]", outMat.Empty())
	compute.SaveMat(*out, outMat)
}
