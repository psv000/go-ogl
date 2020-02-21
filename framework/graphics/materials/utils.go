package materials

import (
	"image"
	"image/png"
	"os"

	"github.com/pkg/errors"
)

func loadJPEG(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Errorf("%s file cant be opened with err %v", filepath, err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, errors.Errorf("%err while decode %s: %v", filepath, err)
	}
	return img, nil
}

func loadPng(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close() // nolint: errcheck
	return png.Decode(f)
}
