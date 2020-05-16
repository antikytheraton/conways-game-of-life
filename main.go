package main

import (
	"runtime"

	"github/antikytheraton/conways-game-of-life/cell"
	"github/antikytheraton/conways-game-of-life/draw"
	"github/antikytheraton/conways-game-of-life/graphic"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 500
	height = 500
)

var (
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
	rows    = 10
	columns = 10
)

func main() {
	runtime.LockOSThread()

	window := graphic.InitGlfw(width, height)
	defer glfw.Terminate()
	program := graphic.InitOpenGL()

	cells := makeCells()
	for !window.ShouldClose() {
		draw.Draw(cells, square, window, program)
	}
}

func makeCells() [][]*cell.Cell {
	cells := make([][]*cell.Cell, rows, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := newCell(x, y)
			cells[x] = append(cells[x], c)
		}
	}

	return cells
}

func newCell(x, y int) *cell.Cell {
	points := make([]float32, len(square), len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(rows)
			position = float32(y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	return &cell.Cell{
		Drawable: draw.MakeVao(points),

		X: x,
		Y: y,
	}
}
