package imageprocessor

type ImageProcessor interface {
	Process(path string) (*string, error)
}
