package detector

import (
	"github.com/golang/mock/gomock"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDetectorsCollection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	obj := NewMockObjectDetectorInterface(ctrl)

	var coll *DetectorsCollection
	Convey("Create a new collection", t, func() {
		coll = NewDetectorsCollection()
		Convey("Add an object to collection", func() {
			coll.Add(obj)
			Convey("The size of collection should be 1", func() {
				So(len(coll.Get()), ShouldEqual, 1)
			})
		})
	})
}
