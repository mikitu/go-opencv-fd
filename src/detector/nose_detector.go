package detector

type NoseDetector struct {
	ObjectDetector
}

func NewNoseDetector() *FaceDetector{
	return &FaceDetector{
		ObjectDetector{Alg : LoadAlgorithm("haarcascade_mcs_nose.xml")},
	}
}



