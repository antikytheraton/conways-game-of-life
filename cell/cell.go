package cell

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// Cell structure for saving matrix cell grid element
type Cell struct {
	Drawable uint32

	X int
	Y int
}

// Draw function to self draw individual cell instance
func (c *Cell) Draw(square []float32) {
	gl.BindVertexArray(c.Drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}
