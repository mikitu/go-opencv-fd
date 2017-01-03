package detector

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/lazywei/go-opencv/opencv"
)

func TestInitializeFaceDetector(t *testing.T) {
	Convey("Create a new object", t, func() {
		fd := NewFaceDetector()
		Convey("Loaded algorithm should be opencv.HaarCascade", func() {
			So(fd.Alg, ShouldHaveSameTypeAs, new(opencv.HaarCascade))
		})
	})
}
