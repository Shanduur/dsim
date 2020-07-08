package compute

import (
	"fmt"
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"image"
	"image/color"
	"unsafe"
)

func Videostream() {
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

func Example() {
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