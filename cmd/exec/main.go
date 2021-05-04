package main

import (
	"flag"

	"github.com/Sheerley/dsim/codes"
	"github.com/Sheerley/dsim/compute"
	"github.com/Sheerley/dsim/fuse"
	"github.com/Sheerley/dsim/plog"
)

var argv fuse.FlagsArray

func main() {
	plog.SetLogLevel(plog.INFO)
	flag.Var(&argv, "img", "names of image file, variable arguments")
	out := flag.String("out", "", "name of output file")
	logLevel := flag.Int("log-level", plog.INFO, "log level")

	flag.Parse()
	if len(argv) < 2 {
		plog.Fatalf(codes.IncorrectArgs, "error with parsing arguments")
	} else if len(argv) > 2 {
		plog.Warningf("too many arguments, skipping %v args", len(argv)-2)
	}
	for _, a := range argv {
		if len(a) == 0 {
			plog.Fatalf(codes.IncorrectArgs, "error with parsing arguments")
		}
	}

	plog.SetLogLevel(*logLevel)

	plog.Messagef("running with args:\n -img1=%v,\n -img2=%v,\n -out=%v,\n -log-level=%v", argv[0], argv[1], *out, *logLevel)

	outMat := compute.BestMatches(argv[0], argv[1])

	plog.Debugf("saving output [empty=%v]", outMat.Empty())
	compute.SaveMat(*out, outMat)
}
