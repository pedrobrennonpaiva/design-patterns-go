package imageprocessor

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type ResizeImageProcessor struct {
	processor ImageProcessor
	width     int
	height    int
}

func NewResizeImageProcessor(processor ImageProcessor, width, height int) *ResizeImageProcessor {
	return &ResizeImageProcessor{
		processor: processor,
		width:     width,
		height:    height,
	}
}

func (p *ResizeImageProcessor) Process(path string) (*string, error) {

	processedImagePath, err := p.processor.Process(path)
	if err != nil {
		return nil, err
	}

	// logic to resize the image
	processedImagePathWithResize := *processedImagePath + fmt.Sprint(p.width+p.height)
	if processedImagePathWithResize == "" {
		return nil, fs.ErrInvalid
	}

	uploadsDir := "structural/decorator/uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		return nil, err
	}

	resizeImagePath := filepath.Join(uploadsDir, "resized_file.jpg")
	err = os.WriteFile(resizeImagePath, []byte(""), 0644)
	if err != nil {
		return nil, err
	}

	return &resizeImagePath, nil
}
