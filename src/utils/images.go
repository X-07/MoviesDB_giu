package utils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

// LoadImage loads image from file and returns *image.RGBA.
func LoadImage(imgPath string) (image.Image, error) {
	imgFile, err := os.Open(filepath.Clean(imgPath))
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error opening image file %s: %w", imgPath, err)
	}

	defer func() {
		// nolint:govet // we want to reuse this err variable here
		if err := imgFile.Close(); err != nil {
			panic(fmt.Sprintf("error closing image file: %s", imgPath))
		}
	}()

	imageData, imageType, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err.Error(), imageType)
		return nil, fmt.Errorf("LoadImage: error decoding jpg/png image: %w", err)
	}

	// img, err := png.Decode(imgFile)
	// if err != nil {
	// 	img, err = jpeg.Decode(imgFile)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return nil, fmt.Errorf("LoadImage: error decoding jpg/png image: %w", err)
	// 	}
	// }

	return imageData, nil
}
