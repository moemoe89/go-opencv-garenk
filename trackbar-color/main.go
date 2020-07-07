package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {

	deviceID := 0

	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Original Video")
	defer window.Close()

	window2 := gocv.NewWindow("Mask Video")
	defer window2.Close()

	window3 := gocv.NewWindow("Result Video")
	defer window3.Close()

	window4 := gocv.NewWindow("Trackbars")
	defer window4.Close()

	img := gocv.NewMat()
	defer img.Close()

	hsv := gocv.NewMat()
	defer hsv.Close()

	mask := gocv.NewMat()
	defer mask.Close()

	result := gocv.NewMat()
	defer result.Close()

	lh := window4.CreateTrackbar("L – H", 179)
	ls := window4.CreateTrackbar("L – S", 255)
	lv := window4.CreateTrackbar("L – V", 255)
	uh := window4.CreateTrackbar("U – H", 179)
	us := window4.CreateTrackbar("U – S", 255)
	uv := window4.CreateTrackbar("U – V", 255)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		gocv.Flip(img, &img, 1)
		gocv.CvtColor(img, &hsv, gocv.ColorBGRToHSV)

		lower := gocv.NewMatFromScalar(gocv.NewScalar(float64(lh.GetPos()), float64(ls.GetPos()), float64(lv.GetPos()), 0.0), gocv.MatTypeCV8U)
		defer lower.Close()
		upper := gocv.NewMatFromScalar(gocv.NewScalar(float64(uh.GetPos()), float64(us.GetPos()), float64(uv.GetPos()), 0.0), gocv.MatTypeCV8U)
		defer upper.Close()

		gocv.InRange(hsv, lower, upper, &mask)
		gocv.BitwiseAndWithMask(img, img, &result, mask)

		window.IMShow(img)
		window2.IMShow(mask)
		window3.IMShow(result)

		if (window.WaitKey(1) >= 27) || (window2.WaitKey(1) >= 27) || (window3.WaitKey(1) >= 27) {
			break
		}

	}

}
