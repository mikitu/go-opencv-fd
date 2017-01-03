package detector

type EyesDetector struct {
	ObjectDetector
}

func NewEyesDetector() *EyesDetector{
	return &EyesDetector{
		ObjectDetector{Alg : LoadAlgorithm("haarcascade_eye_tree_eyeglasses.xml")},
	}
}


