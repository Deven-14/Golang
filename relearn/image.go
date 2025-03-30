package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	rectangle image.Rectangle
	rgba      *image.RGBA
}

func (i Image) Bounds() image.Rectangle {
	return i.rectangle
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return i.rgba.At(x, y)
}

func image_print() {
	rect := image.Rect(0, 0, 100, 100)
	rgba := image.NewRGBA(rect)

	// Initialize pixel data
	for y := range 100 {
		for x := range 100 {
			// Example: Set pixel color based on x and y
			rgba.Set(x, y, color.RGBA{uint8(x * y % 255), uint8(x % 255), uint8(y % 255), 255})
		}
	}

	m := Image{rect, rgba}
	pic.ShowImage(m)
}
