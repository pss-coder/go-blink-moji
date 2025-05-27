package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Hello, 👁️ Blink  😀 Moji  !, press 'q' to exit ")

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

	// Load classifier for face detection
	faceClassifier := gocv.NewCascadeClassifier()
	defer faceClassifier.Close()

	if !faceClassifier.Load("data/face-detection-model/haarcascade_frontalface_default.xml") {
		log.Fatalf("Error loading cascade file: haarcascade_frontalface_default.xml\n")
		return
	}

	// color for the rectangle around the face
	color := color.RGBA{R: 0, G: 255, B: 0, A: 0} // green color

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Error reading from webcam")
			break
		}

		// start detecting faces
		rects := faceClassifier.DetectMultiScale(img)

		// we get the rectangles around the detected faces
		for _, r := range rects {
			gocv.Rectangle(&img, r, color, 3)
			gocv.PutText(&img, "Face", image.Pt(r.Min.X, r.Min.Y-10),
				gocv.FontHersheyPlain, 1.2, color, 2)
		}

		window.IMShow(img)
		window.ResizeWindow(320, 180)

		if window.WaitKey(1) == 113 { // 113 is the ASCII code for 'q'
			fmt.Println("Exiting...")
			break // exit on any key press
		}
	}
}
