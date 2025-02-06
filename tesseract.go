package recognition

import (
	"errors"

	"github.com/otiai10/gosseract/v2"
)

type tesseract struct {
	client *gosseract.Client
}

func NewTesseract(tessdata string) Client {

	c := &tesseract{
		client: gosseract.NewClient(),
	}

	c.client.TessdataPrefix = tessdata

	return c
}

func (t *tesseract) Recognize(dataset string, image []byte) (string, error) {
	if (dataset == "") || (image == nil) {
		return "", errors.New("Dataset or image cannot be empty or nil")
	}

	t.client.SetLanguage(dataset)
	t.client.SetImageFromBytes(image)
	res, err := t.client.Text()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (t *tesseract) Stop() {
	t.client.Close()
}
