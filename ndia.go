package ndia

import (
	"io"

	svg "github.com/ajstarks/svgo"
)

var defautLineHeight = 20

type Canvas struct {
	W, H int

	objects []Object
}

func NewCanvas(w, h int) *Canvas {
	return &Canvas{W: w, H: h}
}

func (c *Canvas) Draw(w io.Writer) {
	canvas := svg.New(w)

	canvas.Start(c.W, c.H)
	// canvas.Startraw(`width="100%"`, `height="100%"`, `viewBox="0 0 600 400"`)
	// canvas.

	for _, o := range c.objects {
		o.draw(canvas, 0, 0)
	}

	canvas.End()
}

func (c *Canvas) AddObject(objects ...Object) {
	for _, o := range objects {
		c.objects = append(c.objects, o)
	}
}

type Align int

const (
	Left   Align = 0
	Center Align = 1
	Right  Align = 2
)

//
// Objects
//

type Object interface {
	draw(*svg.SVG, int, int)
	height() int
	center() (int, int)
}

//Box object
type Box struct {
	X, Y, W, H int
	Style      string

	LineHeight int
	Align      Align

	objects []Object
}

func NewBox(x, y, w, h int, style string) *Box {
	return &Box{X: x, Y: y, W: w, H: h, Style: style, LineHeight: defautLineHeight}
}

func (b *Box) AddObject(objects ...Object) {
	for _, o := range objects {
		b.objects = append(b.objects, o)
	}
}

func (b *Box) draw(canvas *svg.SVG, x, y int) {
	canvas.Rect(b.X+x, b.Y+y, b.W, b.H, b.Style)

	for _, o := range b.objects {

		switch b.Align {
		case Left:
			o.draw(canvas, b.X+x, b.Y+y)
		case Center:
			o.draw(canvas, b.X+b.W/2+x, b.Y+y)
		case Right:
			o.draw(canvas, b.X+b.W+x, b.Y+y)
		}

		h := o.height()
		if h == 0 {
			y += b.LineHeight
		} else {
			y += h
		}
	}
}

func (b *Box) height() int {
	return b.H
}

func (b *Box) center() (int, int) {
	return b.X + b.W/2, b.Y + b.H/2
}

//Circle
type Circle struct {
	X, Y, R int
	Style   string

	Align      Align
	LineHeight int

	objects []Object
}

func NewCircle(x, y, r int, style string) *Circle {
	return &Circle{X: x, Y: y, R: r, Style: style, LineHeight: defautLineHeight}
}

func (c *Circle) AddObject(objects ...Object) {
	for _, o := range objects {
		c.objects = append(c.objects, o)
	}
}

func (c *Circle) draw(canvas *svg.SVG, x, y int) {
	canvas.Circle(c.X, c.Y, c.R, c.Style)

	if c.Align < Right {
		y += c.R
	}

	for _, o := range c.objects {

		switch c.Align {
		case Left:
			o.draw(canvas, c.X-c.R+x, c.Y+y)
		case Center:
			o.draw(canvas, c.X+x, c.Y+y)
		case Right:
			o.draw(canvas, c.X+c.R+x, c.Y+y)
		}

		h := o.height()
		if h == 0 {
			y += c.LineHeight
		} else {
			y += h
		}

	}
}

func (c *Circle) height() int {
	return c.R
}

func (c *Circle) center() (int, int) {
	return c.X, c.Y
}

//Text
type Text struct {
	X, Y, H int
	Text    string
	Style   string
}

func NewText(x, y, h int, title string, style string) *Text {
	return &Text{Text: title, X: x, Y: y, H: h, Style: style}
}

func (t *Text) draw(canvas *svg.SVG, x, y int) {
	canvas.Text(x+t.X, y+t.Y+t.H, t.Text, t.Style)
}

func (t *Text) height() int {
	return t.H + t.Y
}

func (t *Text) center() (int, int) {
	return t.X, t.Y + t.H/2
}

//Line
// type Line struct {
// 	X0, Y0, X1, Y1 int
// 	Style          []string
// }

// func NewLine(x0, y0, x1, y1 int, style ...string) *Line {
// 	return &Line{X0: x0, Y0: y0, X1: x1, Y1: y1, Style: style}
// }

type ConnectedLine struct {
	A, B  Object
	Style string
}

func NewConnectedLine(a, b Object, style string) *ConnectedLine {
	return &ConnectedLine{a, b, style}
}

func (l *ConnectedLine) draw(canvas *svg.SVG, x, y int) {
	x1, y1 := l.A.center()
	x2, y2 := l.B.center()
	canvas.Line(x1, y1, x2, y2, l.Style)
}

func (l *ConnectedLine) height() int {
	return 0
}
func (l *ConnectedLine) center() (int, int) {
	x1, y1 := l.A.center()
	x2, y2 := l.B.center()
	return (x1 + x2) / 2, (y1 + y2) / 2
}
