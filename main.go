package main

import (
	"fmt"
	"image/color"
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

	img := gocv.NewMat()
	defer img.Close()

	eyeClassifier := gocv.NewCascadeClassifier()
	if !eyeClassifier.Load("data/eye-detection-model/haarcascade_eye_tree_eyeglasses.xml") {
		log.Fatalf("Error loading eye classifier: %v\n", err)
		return
	}

	// color for the rectangle around the eye
	color := color.RGBA{R: 0, G: 255, B: 0, A: 0} // green color

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Error reading from webcam")
			break
		}

		rects := eyeClassifier.DetectMultiScale(img)
		// we get the rectangles around the detected faces
		for _, r := range rects {
			gocv.Rectangle(&img, r, color, 3)
		}

		window.IMShow(img)
		window.ResizeWindow(320, 180)

		if window.WaitKey(1) == 113 { // 113 is the ASCII code for 'q'
			fmt.Println("Exiting...")
			break // exit on any key press
		}
	}
}
