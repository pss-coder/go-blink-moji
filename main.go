package main

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Hello, 👁️ Blink  😀 Moji  !")

	// open webcam
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("Error opening video capture device: %v\n", err)
		return
	}
	defer webcam.Close()

	// Create a window to display the video
	window := gocv.NewWindow("Hello Blink Moji")
	defer window.Close()

	// create image matrix to hold the video frame
	img := gocv.NewMat()
	defer img.Close()
	for {
		if ok := webcam.Read(&img); !ok {
			log.Println("Error reading from webcam")
			return
		}
		if img.Empty() {
			log.Println("No image captured from webcam")
			continue
		}

		window.IMShow(img) // display the image in the window

		if window.WaitKey(1) >= 0 {
			fmt.Println("Exiting...")
			break
		}
	}
}
