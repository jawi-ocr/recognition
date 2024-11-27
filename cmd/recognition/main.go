package main

import (
	"fmt"

	"github.com/jawi-ocr/recognition"

	"github.com/mushoffa/gorengan/image"
)

func main() {
	filename := "or_16214_p.2_line_01.png"
	threshold := uint8(180)

	input, err := image.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	m := image.New(input).Monochrome(threshold, false)
	res := recognition.New(m)
	fmt.Println("Recognition: ", res)
}
