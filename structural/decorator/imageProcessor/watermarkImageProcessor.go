package imageprocessor

import (
	"io/fs"
	"os"
	"path/filepath"
)

type WatermarkImageProcessor struct {
	processor ImageProcessor
	watermark string
}

func NewWatermarkImageProcessor(processor ImageProcessor, watermark string) *WatermarkImageProcessor {
	return &WatermarkImageProcessor{
		processor: processor,
		watermark: watermark,
	}
}

func (p *WatermarkImageProcessor) Process(path string) (*string, error) {

	processedImagePath, err := p.processor.Process(path)
	if err != nil {
		return nil, err
	}

	// logic to add watermark to the image
	processedImagePathWithWatermark := *processedImagePath + p.watermark
	if processedImagePathWithWatermark == "" {
		return nil, fs.ErrInvalid
	}

	uploadsDir := "structural/decorator/uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		return nil, err
	}

	watermarkedImagePath := filepath.Join(uploadsDir, "watermarked_file.jpg")
	err = os.WriteFile(watermarkedImagePath, []byte(""), 0644)
	if err != nil {
		return nil, err
	}

	return &watermarkedImagePath, nil
}
