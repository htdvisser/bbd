// Package bbd renders blurred image borders
package bbd

import (
	"image"
	"image/draw"

	"github.com/nfnt/resize"
)

// Render the source image and a blurred background to the destination
func Render(dst draw.Image, src image.Image, details uint) {
	scaleX := float64(dst.Bounds().Dx()) / float64(src.Bounds().Dx())
	scaleY := float64(dst.Bounds().Dy()) / float64(src.Bounds().Dy())
	if scaleX > scaleY {
		scaleX = scaleY
	} else {
		scaleY = scaleX
	}
	sizeX := uint(0.5 + float64(src.Bounds().Dx())*scaleX)
	sizeY := uint(0.5 + float64(src.Bounds().Dy())*scaleY)
	background := resize.Resize(uint(dst.Bounds().Dx()), uint(dst.Bounds().Dy()), resize.Resize(details, details, src, resize.Lanczos3), resize.Lanczos3)
	foreground := resize.Resize(sizeX, sizeY, src, resize.Lanczos3)
	fgPoint := dst.Bounds().Size().Sub(image.Pt(int(sizeX), int(sizeY))).Div(2)
	fgBounds := foreground.Bounds().Add(fgPoint)
	draw.Draw(dst, dst.Bounds(), background, image.ZP, draw.Src)
	draw.Draw(dst, fgBounds, foreground, image.ZP, draw.Src)
}
