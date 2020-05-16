package cell

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// Cell structure for saving matrix cell grid element
type Cell struct {
	Drawable uint32

	alive     bool
	aliveNext bool

	X int
	Y int
}

// Draw function to self draw individual cell instance
func (c *Cell) Draw(square []float32) {
	gl.BindVertexArray(c.Drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

// CheckState determines the state of the cell for the next tick of the game
func (c *Cell) CheckState(cells [][]*Cell) {
	c.alive = c.aliveNext
	c.aliveNext = c.alive

	liveCount := c.LiveNeighbours(cells)
	if c.alive {
		// 1. Any live cell with fewer than two live neighbours dies,
		// as if caused by underpopulation.
		if liveCount < 2 {
			c.aliveNext = false
		}

		// 2. Any live cell with two or three live neighbours lives on
		// the next generation.
		if liveCount == 2 || liveCount == 3 {
			c.aliveNext = true
		}

		// 3. Any live cell with more than three live neighbours dies,
		// as if by overpopulation.
		if liveCount > 3 {
			c.aliveNext = false
		}
	} else {
		// 4. Any dead cell with exactly three live neighbours becomes a
		// live cell, as if by reproduction.
		if liveCount == 3 {
			c.aliveNext = true
		}
	}
}

// LiveNeighbours returns the number of live neighbours for the cell
func (c *Cell) LiveNeighbours(cells [][]*Cell) int {
	var liveCount int
	add := func(x, y int) {
		// If we're at the edge, check the other side of the board
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].alive {
			liveCount++
		}
	}

	add(c.X-1, c.Y)   // to the left
	add(c.X+1, c.Y)   // to the right
	add(c.X, c.Y+1)   // up
	add(c.X, c.Y-1)   // down
	add(c.X-1, c.Y+1) // top-left
	add(c.X+1, c.Y+1) // top-right
	add(c.X-1, c.Y-1) // bottom-left
	add(c.X+1, c.Y-1) // bottom-right

	return liveCount
}
