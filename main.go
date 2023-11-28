package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Pixel struct {
	color uint32
}

func (p *Pixel) RGBA() (r, g, b, a uint32) {
	red := ((p.color >> (8 * 3)) & 0xff) << 8   // red
	green := ((p.color >> (8 * 2)) & 0xff) << 8 // green
	blue := ((p.color >> (8 * 1)) & 0xff) << 8  // blue
	alpha := ((p.color >> (8 * 0)) & 0xff) << 8

	return red * alpha, green * alpha, blue * alpha, alpha
}

type Canvas struct {
	width  int
	height int
	pixels []Pixel
}

func (c *Canvas) Opaque() bool {

	b := c.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			_, _, _, a := c.At(x, y).RGBA()
			if a != 0xff00 {
				return false
			}
		}
	}
	return true
}

func (c *Canvas) ColorModel() color.Model {
	return color.RGBAModel
}

func (c *Canvas) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.width, c.height)
}

func (c *Canvas) At(x, y int) color.Color {
	return &c.pixels[y*c.width+x]
}

func NewCanvas(width int, height int, c uint32) *Canvas {
	pixels := make([]Pixel, width*height)
	pixel := Pixel{
		color: c,
	}

	for i := range pixels {
		pixels[i] = pixel
	}

	return &Canvas{
		width:  width,
		height: height,
		pixels: pixels,
	}
}

var (
	WIDTH             = 800
	HEIGHT            = 600
	BACKGROUND uint32 = 0xFFFF00FF
)

func main() {

	canvas := NewCanvas(WIDTH, HEIGHT, BACKGROUND)

	f, err := os.Create("image.png")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, canvas); err != nil {
		if err != nil {
			return
		}
		log.Fatal(err)
	}

}
