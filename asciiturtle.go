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
	penUp   bool
}

func NewPen(canvas Canvas, char byte, x, y int) (*Pen, error) {
	if canvas == nil {
		return nil, fmt.Errorf("canvas must not be nil")
	}

	switch {
	case x < 0:
		x = 0
	case x >= canvas.Width():
		x = canvas.Width() - 1
	}

	switch {
	case y < 0:
		y = 0
	case y >= canvas.Height():
		y = canvas.Height() - 1
	}

	return &Pen{
		Canvas: canvas,
		X:      x,
		Y:      y,
		Char:   char,
	}, nil
}

func (p *Pen) PenUp() {
	p.penUp = true
}

func (p *Pen) PenDown() {
	p.penUp = false
}

func (p *Pen) Dot() {
	if p.penUp {
		return
	}

	p.Canvas[p.Y][p.X] = p.Char
}

func (p *Pen) Goto(x, y int) {
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

	p.X = x
	p.Y = y
}

func (p *Pen) Forward(distance int) {
	d := float64(distance)

	x1 := p.X + int(d*math.Cos(p.Heading*degToRad))
	y1 := p.Y - int(d*math.Sin(p.Heading*degToRad))

	p.drawTo(x1, y1)

}

func (p *Pen) Backward(distance int) {
	d := float64(distance)

	x1 := p.X - int(d*math.Cos(p.Heading*degToRad))
	y1 := p.Y + int(d*math.Sin(p.Heading*degToRad))

	p.drawTo(x1, y1)
}

func (p *Pen) drawTo(x1, y1 int) {
	x0 := p.X
	y0 := p.Y

	d := math.Max(math.Abs(float64(x1)-float64(x0)),
		math.Abs(float64(y1)-float64(y0)))

	for i := 0; i <= int(d); i++ {
		t := float64(i) / d
		x, y := lerpPoint(x0, y0, x1, y1, t)

		if x1 < 0 || x1 >= p.Canvas.Width() ||
			y1 < 0 || y1 >= p.Canvas.Height() {
			break
		}

		p.X = x
		p.Y = y

		p.Dot()
	}
}

func lerpPoint(x0, y0, x1, y1 int, t float64) (int, int) {
	x := int(math.Round(lerp(float64(x0), float64(x1), t)))
	y := int(math.Round(lerp(float64(y0), float64(y1), t)))

	return x, y
}

func lerp(n, m, t float64) float64 {
	return n + t*(m-n)
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
