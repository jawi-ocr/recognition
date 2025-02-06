package main

import (
	"fmt"

	"github.com/jawi-ocr/recognition"

	"github.com/mushoffa/gorengan/image"
)

func main() {
	filename := "OR_16214_P001L0101_W_202.png"
	tessdata := "/media/mushoffa/DATA/4_WORK/2_UTM/1_JAWI_OCR/2_CODE/jawi-ocr/recognition/model/"
	// threshold := uint8(180)

	input, err := image.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// m := image.New(input).Monochrome(threshold, false)
	m := image.New(input)
	img, err := m.Bytes()
	if err != nil {
		panic(err)
	}

	client := recognition.NewTesseract(tessdata)
	res, err := client.Recognize("jawi", img)
	if err != nil {
		panic(err)
	}
	// res := recognition.NewWithDataset(input, "jawi")
	fmt.Println("Recognition: ", res)
}
