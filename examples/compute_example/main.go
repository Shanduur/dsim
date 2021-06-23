package main

import (
	"os"

	"github.com/Sheerley/dsip/codes"
	"github.com/Sheerley/dsip/compute"
	"github.com/Sheerley/dsip/plog"
)

func main() {
	//compute.Videostream()
	//fmt.Println(gocv.Version(), gocv.OpenCVVersion())
	res := compute.BestMatches(
		"/home/shanduur/repos/dsip/examples/images/box.png",
		"/home/shanduur/repos/dsip/examples/images/box_in_scene.png")

	if res.Rows() > 0 && res.Cols() > 0 {
		plog.Messagef("git gud")
		os.Exit(0)
	}
	plog.Fatalf(codes.ComputeError, "improper size of res matrix")
}
