package detector
import(
	"github.com/lazywei/go-opencv/opencv"
	"fmt"
	"os"
	"path"
)
type DetectorExecutor struct {
	win  *opencv.Window
	cap *opencv.Capture
}

func NewDetectorExecutor(win  *opencv.Window) *DetectorExecutor {
	capt := opencv.NewCameraCapture(0)
	if capt == nil {
		panic("cannot open camera")
	}
	return &DetectorExecutor{win: win, cap: capt}
}

func (d DetectorExecutor) Run(detectors []ObjectDetectorInterface) {
	for {
		img := d.retreieveFrame()
		for _, obj := range detectors {
			obj.Detect(img)
		}
		d.win.ShowImage(img)

		key := opencv.WaitKey(1)

		if key == 27 {
			os.Exit(0)
		}
	}

}

func (d DetectorExecutor) retreieveFrame() *opencv.IplImage {
	if ! d.cap.GrabFrame() {
		fmt.Println("cannot get frame")
		os.Exit(0)
	}
	img := d.cap.RetrieveFrame(1)
	if img == nil {
		fmt.Println("nil image")
		os.Exit(0)
	}
	return img
}

func (d DetectorExecutor) LoadFile(file string) *opencv.HaarCascade {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return opencv.LoadHaarClassifierCascade(path.Join(cwd, "algorithms/haarcascades/" + file))
}

func (d DetectorExecutor) Release() {
	d.cap.Release()
}

