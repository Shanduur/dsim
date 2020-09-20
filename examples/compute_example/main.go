package main

import (
	"github.com/Sheerley/pluggabl/pkg/compute"
	"os"
)

func main() {
	//compute.Videostream()
	//fmt.Println(gocv.Version(), gocv.OpenCVVersion())
	res := compute.BestMatches("/home/shanduur/Pictures/osd1.jpg", "/home/shanduur/Pictures/osd.jpg")

	if res.Rows() > 0 && res.Cols() > 0 {
		os.Exit(0)
	}
	os.Exit(1)
}