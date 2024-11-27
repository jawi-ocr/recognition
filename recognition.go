package recognition

import (
	"bytes"
	"image"
	"image/png"
)

func New(img image.Image) string {
	data := encode(img)
	return tesseract(data, "ara")
}

func NewWithDataset(img image.Image, dataset string) string {
	data := encode(img)
	return tesseract(data, dataset)
}

func encode(img image.Image) []byte {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		return nil
	}

	return buffer.Bytes()
}
