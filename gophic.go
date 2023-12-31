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
	r = (p.color >> (8 * 3)) & 0xFF
	g = (p.color >> (8 * 2)) & 0xFF
	b = (p.color >> (8 * 1)) & 0xFF
	a = (p.color >> (8 * 0)) & 0xFF

	r = uint32(r) * 0x101
	g = uint32(g) * 0x101
	b = uint32(b) * 0x101
	a = uint32(a) * 0x101
	return r * a / 0xffff, g * a / 0xfff, b * a / 0xffff, a
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

//func (c *Canvas) Opaque() bool {
//	b := c.Bounds()
//	for y := b.Min.Y; y < b.Max.Y; y++ {
//		for x := b.Min.X; x < b.Max.X; x++ {
//			_, _, _, a := c.At(x, y).RGBA()
//			if a != 0xffff {
//				return false
//			}
//		}
//	}
//	return true
//}

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

func fillRect(canvas *Canvas, x0 int, y0 int, w uint, h uint, color uint32) {
	pixel := Pixel{
		color: color,
	}

	for dy := 0; dy < int(h); dy++ {
		y := y0 + dy
		if y >= 0 && y < canvas.height {
			for dx := 0; dx < int(w); dx++ {
				x := x0 + dx
				if x >= 0 && x < canvas.width {
					canvas.pixels[y*canvas.width+x] = pixel
				}
			}
		}
	}
}

func fillCircle(canvas *Canvas, cx int, cy int, r uint, color uint32) {
	x1 := cx - int(r)
	y1 := cy - int(r)
	x2 := cx + int(r)
	y2 := cy + int(r)

	pixel := Pixel{
		color: color,
	}

	for y := y1; y <= y2; y++ {
		if y >= 0 && y < canvas.height {
			for x := x1; x <= x2; x++ {
				if x >= 0 && x < canvas.width {
					dx := x - cx
					dy := y - cy

					if uint(dx*dx+dy*dy) <= r*r {
						canvas.pixels[y*canvas.width+x] = pixel
					}
				}
			}
		}
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func drawLine(canvas *Canvas, x1 int, y1 int, x2 int, y2 int, color uint32) {
	dx := x2 - x1
	dy := y2 - y1

	pixel := Pixel{
		color: color,
	}
	if dx != 0 {
		c := y1 - dy*x1/dx
		if x1 > x2 {
			swap(&x1, &x2)
		}
		for x := x1; x <= x2; x++ {
			if x >= 0 && x < canvas.width {
				sy1 := dy*x/dx + c
				sy2 := dy*(x+1)/dx + c
				if sy1 > sy2 {
					swap(&sy1, &sy2)
				}
				for y := sy1; y <= sy2; y++ {
					if y >= 0 && y < canvas.height {
						canvas.pixels[y*canvas.width+x] = pixel
					}
				}
			}
		}
	} else {
		if y1 > y2 {
			swap(&y1, &y2)
		}
		for y := y1; y <= y2; y++ {
			if y >= 0 && y < canvas.height {
				canvas.pixels[y*canvas.width+x1] = pixel

			}
		}
	}

}

func saveToPNG(canvas *Canvas, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("ERROR: unable to close file: %v", err)
		}
	}(f)

	if err := png.Encode(f, canvas); err != nil {
		log.Fatalf("ERROR: unable to encode file to PNG: %v", err)
	}
}
