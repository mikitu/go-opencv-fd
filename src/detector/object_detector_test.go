package detector

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/lazywei/go-opencv/opencv"
	"github.com/golang/mock/gomock"
)

func TestDetect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//img := NewMockIplImageInterface(ctrl)

	Convey("Create a new object", t, func() {
		fd := new(ObjectDetector)
		fd.Alg = LoadAlgorithm("haarcascade_frontalface_alt.xml")
		Convey("Loaded algorithm should be opencv.HaarCascade type", func() {
			So(fd.Alg, ShouldHaveSameTypeAs, new(opencv.HaarCascade))
		})
		//fd.Detect(img)

	})
}
