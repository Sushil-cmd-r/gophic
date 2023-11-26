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
	r = (p.color >> (8 * 0)) & 0xFF
	g = (p.color >> (8 * 1)) & 0xFF
	b = (p.color >> (8 * 2)) & 0xFF
	a = (p.color >> (8 * 3)) & 0xFF

	return r, g, b, a
}

type Canvas struct {
	width  int
	height int
	pixels []Pixel
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

func (c *Canvas) Opaque() bool {
	b := c.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			_, _, _, a := c.At(x, y).RGBA()
			if a != 0xff {
				return false
			}
		}
	}
	return true
}

func NewCanvas(width int, height int, color uint32) *Canvas {
	pixels := make([]Pixel, width*height)
	pixel := Pixel{
		color: color,
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
	BACKGROUND uint32 = 0xFF0000FF
)

func main() {
	canvas := NewCanvas(WIDTH, HEIGHT, BACKGROUND)
	//const width, height = 256, 256

	// Create a colored image of the given width and height.
	//img := image.NewNRGBA(image.Rect(0, 0, width, height))
	//
	//for y := 0; y < height; y++ {
	//	for x := 0; x < width; x++ {
	//		img.Set(x, y, color.NRGBA{
	//			R: uint8((x + y) & 255),
	//			G: uint8((x + y) << 1 & 255),
	//			B: uint8((x + y) << 2 & 255),
	//			A: 255,
	//		})
	//	}
	//}
	//
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, canvas); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
