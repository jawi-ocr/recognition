package recognition

import (
	"github.com/otiai10/gosseract/v2"
)

func tesseract(data []byte, dataset string) string {
	if (data == nil) || (dataset == "") {
		return ""
	}

	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage(dataset)
	client.SetImageFromBytes(data)
	res, _ := client.Text()

	return res
}
