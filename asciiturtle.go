package asciiturtle

import (
	"fmt"
	"math"
	"strings"
)

const degToRad = math.Pi / 180.0

type Pen struct {
	Canvas  Canvas
	X, Y    int
	Char    byte
	Heading float64
}

func NewPen(canvas Canvas, char byte, x, y int) (*Pen, error) {
	if canvas == nil {
		return nil, fmt.Errorf("canvas must not be nil")
	}

	if x < 0 || x >= canvas.Width() {
		x = 0
	}

	if y < 0 || y >= canvas.Width() {
		y = 0
	}

	return &Pen{
		Canvas: canvas,
		X:      x,
		Y:      y,
		Char:   char,
	}, nil
}

func (p *Pen) Dot() {
	p.Canvas[p.Y][p.X] = p.Char
}

func (p *Pen) Goto(x, y int) {
	if x < 0 || x >= p.Canvas.Width() || y < 0 || y >= p.Canvas.Height() {
		return
	}

	switch {
	case x < 0:
		x = 0
	case x >= p.Canvas.Width():
		x = p.Canvas.Width() - 1
	}

	switch {
	case y < 0:
		y = 0
	case y >= p.Canvas.Height():
		y = p.Canvas.Height() - 1
	}

	for p.X != x || p.Y != y {
		switch {
		case p.X < x:
			p.X++
		case p.X > x:
			p.X--
		}

		switch {
		case p.Y < y:
			p.Y++
		case p.Y > y:
			p.Y--
		}

		p.Dot()
	}
}

func (p *Pen) Forward(distance int) {
	d := float64(distance)

	x := p.X + int(d*math.Cos(p.Heading*degToRad))
	y := p.Y - int(d*math.Sin(p.Heading*degToRad))

	p.Goto(x, y)
}

func (p *Pen) Backward(distance int) {
	d := float64(distance)

	x := p.X - int(d*math.Cos(p.Heading*degToRad))
	y := p.Y + int(d*math.Sin(p.Heading*degToRad))

	p.Goto(x, y)
}

func (p *Pen) Right(deg float64) {
	p.Heading -= deg
}

func (p *Pen) Left(deg float64) {
	p.Heading += deg
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
