package compute

import (
	"image/color"

	"gocv.io/x/gocv"
)

// BestMatches takes two arguments being Querry and Train
// and holding paths to the images that have to be
// processed during the SIFT detection and Brute
// force matching. Function returns gocv.Mat holding
// result of feature matching algorithm.
func BestMatches(query string, train string) gocv.Mat {

	img1 := gocv.IMRead(query, gocv.IMReadGrayScale)
	defer img1.Close()
	img2 := gocv.IMRead(train, gocv.IMReadGrayScale)
	defer img2.Close()

	// creating new SIFT
	sift := gocv.NewSIFT()
	defer sift.Close()

	// detecting and computing keypoints using SIFT method
	kp1, des1 := sift.DetectAndCompute(img1, gocv.NewMat())
	kp2, des2 := sift.DetectAndCompute(img2, gocv.NewMat())

	// finding K best matches for each descriptor
	bf := gocv.NewBFMatcher()
	matches := bf.KnnMatch(des1, des2, 2)

	// application of ratio test
	var good []gocv.DMatch
	for _, m := range matches {
		if len(m) > 1 {
			if m[0].Distance < 0.75*m[1].Distance {
				good = append(good, m[0])
			}
		}
	}

	// matches color
	c1 := color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 0,
	}

	// point color
	c2 := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}

	// creating empty mask
	mask := make([]byte, 0)

	// new matrix for output image
	out := gocv.NewMat()

	// drawing matches
	gocv.DrawMatches(img1, kp1, img2, kp2, good, &out, c1, c2, mask, gocv.DrawDefault)

	return out
}

func SaveMat(name string, img gocv.Mat) {
	gocv.IMWrite(name, img)
}
