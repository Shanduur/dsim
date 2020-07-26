package main

import "github.com/Sheerley/pluggabl/pkg/compute"

func main() {
	//compute.Videostream()
	//fmt.Println(gocv.Version(), gocv.OpenCVVersion())
	compute.BruteSIFTMatching("/home/shanduur/Pictures/osd1.jpg", "/home/shanduur/Pictures/osd.jpg")
}