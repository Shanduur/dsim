package main

import (
	"os"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/pkg/compute"
	"github.com/Sheerley/pluggabl/pkg/plog"
)

func main() {
	//compute.Videostream()
	//fmt.Println(gocv.Version(), gocv.OpenCVVersion())
	res := compute.BestMatches("/home/shanduur/Pictures/osd1.jpg", "/home/shanduur/Pictures/osd.jpg")

	if res.Rows() > 0 && res.Cols() > 0 {
		os.Exit(0)
	}
	plog.Fatalf(codes.ComputeError, "improper size of res matrix")
}
