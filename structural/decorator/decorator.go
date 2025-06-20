package decorator

import (
	imageprocessor "design-patterns-go/structural/decorator/imageProcessor"
	"fmt"
)

func Run() {

	imagePath := "/tmp/file.jpg"
	imgProcessor := imageprocessor.NewBasicImageProcessor()

	imgProcessor = imageprocessor.NewWatermarkImageProcessor(imgProcessor, " - Watermark")

	imgProcessor = imageprocessor.NewResizeImageProcessor(imgProcessor, 800, 600)

	if _, err := imgProcessor.Process(imagePath); err != nil {
		fmt.Println("Error processing image:", err)
		return
	}

	fmt.Println("Decorator processed successfully")
}
