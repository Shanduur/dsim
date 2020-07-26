package compute

import (
	"fmt"
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"image"
	"image/color"
	"unsafe"
)

func BruteSIFTMatching(query string, train string) {
	img1 := gocv.IMRead(query, gocv.IMReadGrayScale)
	img2 := gocv.IMRead(train, gocv.IMReadGrayScale)

	sift := contrib.NewSIFT()
	defer sift.Close()

	kp1, des1 := sift.DetectAndCompute(img1, gocv.NewMat())
	kp2, des2 := sift.DetectAndCompute(img2, gocv.NewMat())

	bf := gocv.NewBFMatcher()
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

	// Temporary solution until DrawKeyPointsBGRA comes into mainstream
	gocv.CvtColor(img1, &img1, gocv.ColorBGRToRGBA)
	gocv.CvtColor(img2, &img2, gocv.ColorBGRToRGBA)

	window1 := gocv.NewWindow("Query")
	window1.IMShow(img1)

	window2 := gocv.NewWindow("Train")
	window2.IMShow(img2)

	fmt.Println(len(good))

	window1.WaitKey(0)
	window2.WaitKey(0)
}

// Function testing SIFT real time performance using
// video capture device.
func EXAMPLE_VideoStream() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Window 1")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		detector(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}

func detector(image *gocv.Mat) {
	sift := contrib.NewSIFT()

	keypoints := sift.Detect(*image)

	var kp gocv.KeyPoint

	fmt.Println(float64(len(keypoints) * int(unsafe.Sizeof(kp)))/(1024), "kb" )

	gocv.DrawKeyPoints(*image, keypoints, image, color.RGBA{
		B: 255,
		G: 0,
		R: 0,
		A: 0,
	}, gocv.DrawDefault)
}

// Function showing with exemplary use of the SIFT algorithm
// on static images.
func EXAMPLE_StaticImage() {
	window1 := gocv.NewWindow("Example")
	img := gocv.IMRead("/home/shanduur/Pictures/osd.jpg", gocv.IMReadColor)
	sift1 := contrib.NewSIFT()
	kp1 := sift1.Detect(img)
	gocv.DrawKeyPoints(img, kp1, &img, color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 0,
	}, gocv.DrawDefault)

	gray := gocv.NewMat()
	gocv.CvtColor(img, &gray, gocv.ColorRGBToGray)

	sift2 := contrib.NewSIFT()
	kp2 := sift2.Detect(gray)

	gocv.DrawKeyPoints(gray, kp2, &gray, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, gocv.DrawDefault)

	fmt.Printf("Length of color kp list: %v\t Size: %vkb\n", len(kp1), float64(len(kp1)*int(unsafe.Sizeof(kp1[0])))/1024)
	fmt.Printf("Length of gray kp list: %v\t Size: %vkb\n", len(kp2), float64(len(kp2)*int(unsafe.Sizeof(kp2[0])))/1024)

	result := gocv.NewMat()

	gocv.MatchTemplate(img, gray, &result, gocv.TmSqdiffNormed, gocv.NewMat())

	_, _, mnLoc, _ := gocv.MinMaxLoc(result)

	size := img.Size()

	gocv.Rectangle(&gray, image.Rectangle{
		Min: image.Point{
			X: mnLoc.X,
			Y: mnLoc.Y,
		},
		Max: image.Point{
			X: mnLoc.X + size[0],
			Y: mnLoc.Y + size[1],
		},
	}, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, 2)

	window1.IMShow(gray)
	window1.WaitKey(0)
}