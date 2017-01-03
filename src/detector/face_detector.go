package detector

type FaceDetector struct {
	ObjectDetector
}

func NewFaceDetector() *FaceDetector{
	return &FaceDetector{
		ObjectDetector{Alg : LoadAlgorithm("haarcascade_frontalface_alt.xml")},
	}
}


