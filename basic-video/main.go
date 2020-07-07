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

	window := gocv.NewWindow("Basic Video")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		gocv.Flip(img, &img, 1)

		window.IMShow(img)
		if window.WaitKey(1) >= 27 {
			break
		}
	}

}
