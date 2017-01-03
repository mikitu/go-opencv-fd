package main

import (
	"fmt"

	"github.com/lazywei/go-opencv/opencv"
	"github.com/mikitu/go-opencv-fd/src/detector"
)

func main() {

	win := opencv.NewWindow("Go-OpenCV Webcam Face and Eyes Detection")
	defer win.Destroy()

	fmt.Println("Press ESC to quit")

	detectors := detector.NewDetectorsCollection()
	d1:=detector.NewFaceDetector()
	detectors.Add(d1)

	detectors.Add(detector.NewEyesDetector())

	executor := detector.NewDetectorExecutor(win)
	defer executor.Release()

	executor.Run(detectors.Get())

}

