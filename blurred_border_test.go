package bbd_test

import (
	"image"
	"image/jpeg"
	"os"

	"go.htdvisser.nl/bbd"
)

func ExampleRender() {
	// Open the source file
	srcFile, err := os.Open("doc/src.jpg")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	// Read the image from the source file
	src, _, err := image.Decode(srcFile)
	if err != nil {
		panic(err)
	}

	// Create the destination image
	img := image.NewRGBA(image.Rect(0, 0, 1920, 1080))

	// Render
	bbd.Render(img, src, 3)

	// Open the destination file
	f, err := os.OpenFile("doc/dst.jpg", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()
	err = jpeg.Encode(f, img, &jpeg.Options{
		Quality: 90,
	})
	if err != nil {
		panic(err)
	}
}
