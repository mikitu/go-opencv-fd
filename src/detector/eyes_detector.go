package detector

import (
	"github.com/lazywei/go-opencv/opencv"
)

type EyesDetector struct {
	ObjectDetector
}

func NewEyesDetector() *EyesDetector{
	return &EyesDetector{
		ObjectDetector{Alg : LoadAlgorithm("haarcascade_eye_tree_eyeglasses.xml")},
	}
}


func (d EyesDetector) Draw(img IplImageInterface) {
	for _, value := range d.objects {
		opencv.Rectangle(img.(*opencv.IplImage),
			opencv.Point{value.X(), value.Y()},
			opencv.Point{value.X() + value.Width(), value.Y() + value.Height()},
			opencv.ScalarAll(255.0), 1, 1, 0)
	}
}