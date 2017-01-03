package detector

import (
	"github.com/lazywei/go-opencv/opencv"
	"os"
	"path"
	"unsafe"
)

type IplImageInterface interface {
	Channels() int
	Depth() int
	Origin() int
	Width() int
	Height() int
	WidthStep() int
	ImageSize() int
	ImageData() unsafe.Pointer
}

type HaarCascadeInterface interface {
	DetectObjects(img *opencv.IplImage) []*opencv.Rect
	Release()
}
type ObjectDetectorInterface interface {
	Detect(img IplImageInterface)
	Draw(img IplImageInterface)

}

type ObjectDetector struct{
	Alg HaarCascadeInterface
	objects []*opencv.Rect
}

func (d *ObjectDetector) Detect(img IplImageInterface){
	d.objects = d.Alg.DetectObjects(img.(*opencv.IplImage))
}

func (d ObjectDetector) Draw(img IplImageInterface) {
	for _, value := range d.objects {
		opencv.Circle(img.(*opencv.IplImage),
			opencv.Point{
				value.X() + (value.Width() / 2),
				value.Y() + (value.Height() / 2),
			},
			value.Width() / 2,
			opencv.ScalarAll(255.0), 1, 1, 0)
	}
}

func LoadAlgorithm(file string) HaarCascadeInterface {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return opencv.LoadHaarClassifierCascade(path.Join(cwd, "src/algorithms/haarcascades/" + file))
}

