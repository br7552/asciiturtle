package asciiturtle

import (
	"fmt"
	"math"
	"strings"
)

type Pen struct {
	Canvas  Canvas
	X, Y    int
	Char    byte
	Heading int
}

func NewPen(canvas Canvas, char byte, x, y int) (*Pen, error) {
	if canvas == nil {
		return nil, fmt.Errorf("canvas must not be nil")
	}

	pen := Pen{
		Canvas: canvas,
		Char:   char,
	}

	err := pen.Goto(x, y)
	if err != nil {
		return nil, err
	}

	return &pen, nil
}

func (p *Pen) Dot() {
	p.Canvas[p.Y][p.X] = p.Char
}

func (p *Pen) Goto(x, y int) error {
	if x < 0 || x >= p.Canvas.Width() || y < 0 || y >= p.Canvas.Height() {
		return fmt.Errorf("(%d,%d) must be within %dx%d canvas", x, y,
			p.Canvas.Width(), p.Canvas.Height())
	}

	p.X = x
	p.Y = y

	return nil
}

func (p *Pen) Forward(distance int) {
	for i := 0; i < distance; i++ {
		x := p.X + int(math.Round(math.Cos(float64(p.Heading))))
		y := p.Y + int(math.Round(math.Sin(float64(p.Heading))))

		if err := p.Goto(x, y); err != nil {
			break
		}

		p.Dot()
	}
}

func (p *Pen) Backward(distance int) {
	for i := 0; i < distance; i++ {
		x := p.X - int(math.Round(math.Cos(float64(p.Heading))))
		y := p.Y - int(math.Round(math.Sin(float64(p.Heading))))

		if err := p.Goto(x, y); err != nil {
			break
		}

		p.Dot()
	}
}

func (p *Pen) Right(deg int) {
	p.Heading += deg
}

func (p *Pen) Left(deg int) {
	p.Heading -= deg
}

type Canvas [][]byte

func NewCanvas(x, y int) Canvas {
	if x <= 0 || y <= 0 {
		return [][]byte{}
	}

	canvas := make([][]byte, y)

	for i := range canvas {
		canvas[i] = make([]byte, x)
	}

	return canvas
}

func (c Canvas) String() string {
	var b strings.Builder

	for _, row := range c {
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

func (c Canvas) Height() int {
	return len(c)
}

func (c Canvas) Width() int {
	if len(c) == 0 {
		return 0
	}

	return len(c[0])
}
