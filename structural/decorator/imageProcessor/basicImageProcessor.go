package imageprocessor

import (
	"os"
	"path/filepath"
)

type BasicImageProcessor struct{}

func NewBasicImageProcessor() ImageProcessor {
	return &BasicImageProcessor{}
}

func (p *BasicImageProcessor) Process(path string) (*string, error) {
	// logic to do processing of the image

	uploadsDir := "structural/decorator/uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		return nil, err
	}

	newImagePath := filepath.Join(uploadsDir, "file.jpg")
	err := os.WriteFile(newImagePath, []byte(""), 0644)
	if err != nil {
		return nil, err
	}

	return &newImagePath, nil
}
