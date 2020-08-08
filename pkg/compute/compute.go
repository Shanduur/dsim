package compute

import (
	"gocv.io/x/gocv"
	"image/color"
)

func BestMatches(query string, train string) gocv.Mat {

	img1 := gocv.IMRead(query, gocv.IMReadGrayScale)
	defer img1.Close()
	img2 := gocv.IMRead(train, gocv.IMReadGrayScale)
	defer img2.Close()

	sift := gocv.NewSIFT()
	defer sift.Close()

	kp1, des1 := sift.DetectAndCompute(img1, gocv.NewMat())
	kp2, des2 := sift.DetectAndCompute(img2, gocv.NewMat())

	bf := gocv.NewBFMatcher()
	defer bf.Close()
	matches := bf.KnnMatch(des1, des2, 2)

	var good []gocv.DMatch

	for _, m := range matches {
		if len(m) > 1 {
			if m[0].Distance < 0.75 * m[1].Distance {
				good = append(good, m[0])
			}
		}
	}

	gocv.DrawKeyPoints(img1, kp1, &img1, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, gocv.DrawDefault)

	gocv.DrawKeyPoints(img2, kp2, &img2, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, gocv.DrawDefault)

	out := gocv.NewMat()
	defer out.Close()

	//gocv.DrawMatches()

	return out
}