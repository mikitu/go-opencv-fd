package detector

type DetectorsCollection struct {
	objects []ObjectDetectorInterface
}

func NewDetectorsCollection() *DetectorsCollection {
	return &DetectorsCollection{}
}
func (c *DetectorsCollection) Add(obj ObjectDetectorInterface) {
	c.objects = append(c.objects, obj)
}
func (c DetectorsCollection) Get() []ObjectDetectorInterface {
	return c.objects
}