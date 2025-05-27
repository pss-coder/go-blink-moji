package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

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
	// what is a matrix -> digital representation of the image using a grid (or matrix) of pixels.
	// each pixel contains color or brightness information.
	img := gocv.NewMat()
	defer img.Close()

	// let's gray scale our image
	gray := gocv.NewMat()
	defer gray.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			log.Println("Error reading from webcam")
			return
		}
		if img.Empty() {
			log.Println("No image captured from webcam")
			continue
		}
		gocv.Flip(img, &img, 1) // mirror the image horizontally
		gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

		gocv.PutText(&gray, time.Now().Format("2006-01-02 15:04:05"), image.Pt(10, 30), gocv.FontHersheySimplex, 1.0, color.RGBA{0, 255, 0, 0}, 2)

		window.IMShow(gray)           // display the image in the window
		window.ResizeWindow(320, 320) // resize the window to 100x100 pixels

		if window.WaitKey(1) >= 0 {
			fmt.Println("Exiting...")
			break
		}
	}
}
