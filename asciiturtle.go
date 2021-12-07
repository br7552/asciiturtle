package asciiturtle

import "strings"

type Grid [][]byte

func NewGrid(x, y int) Grid {
	if x <= 0 || y <= 0 {
		return [][]byte{}
	}

	grid := make([][]byte, y)

	for i := range grid {
		grid[i] = make([]byte, x)
	}

	return grid
}

func (g Grid) String() string {
	var b strings.Builder

	for _, row := range g {
		for _, v := range row {
			if v == 0 {
				b.WriteByte(' ')
				continue
			}

			b.WriteByte(v)
		}

		b.WriteByte('\n')
	}

	return b.String()
}
