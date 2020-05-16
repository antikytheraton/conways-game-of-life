package main

import (
	"runtime"

	"github/antikytheraton/conways-game-of-life/graphic"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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
	shapeCount = int32(len(square) / 3)

	rows    = 10
	columns = 10
)

type cell struct {
	drawable uint32

	x int
	y int
}

func main() {
	runtime.LockOSThread()

	window := graphic.InitGlfw()
	defer glfw.Terminate()
	program := graphic.InitOpenGL()

	vao := makeVao(square)
	for !window.ShouldClose() {
		draw(vao, shapeCount, window, program)
	}
}

// func makeCells() [][]*cell {
// 	cells := make([][]*cell, rows, rows)
// 	for x := 0; x < rows; x++ {
// 		for y := 0; y < columns; y++ {
// 			c := newCell(x, y)
// 			cells[x] = append(cells[x], c)
// 		}
// 	}
// }

func draw(vao uint32, count int32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, count)

	glfw.PollEvents()
	window.SwapBuffers()
}

// makeVao (Vertex Array Object) initializes and returns a vertex
// array from the points provided
func makeVao(points []float32) uint32 {
	var vbo uint32 // Vertex Buffer Object
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
