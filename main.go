package main

var (
	WIDTH             = 800
	HEIGHT            = 600
	BACKGROUND uint32 = 0xFFFFFFFF
	FOREGROUND uint32 = 0xFF0000FF

	COLS       = 8
	ROWS       = 6
	CellWidth  = uint(WIDTH / COLS)
	CellHeight = uint(HEIGHT / ROWS)
)

func checkers() {
	canvas := NewCanvas(WIDTH, HEIGHT, BACKGROUND)

	for y := 0; y < ROWS; y++ {
		for x := 0; x < COLS; x++ {
			if (x+y)%2 == 0 {
				fillRect(canvas, x*int(CellWidth), y*int(CellHeight), CellWidth, CellHeight, FOREGROUND)
			}
		}
	}
	saveToPNG(canvas, "img/checker.png")

}

func circle() {
	canvas := NewCanvas(WIDTH, HEIGHT, BACKGROUND)

	for y := 0; y < ROWS; y++ {
		for x := 0; x < COLS; x++ {
			r := CellWidth
			if CellHeight < r {
				r = CellHeight
			}
			fillCircle(canvas, x*int(CellWidth)+int(CellWidth/2), y*int(CellHeight)+int(CellHeight/2), r/2, FOREGROUND)
		}
	}
	saveToPNG(canvas, "img/circle.png")
}

func main() {
	checkers()
	circle()
}
