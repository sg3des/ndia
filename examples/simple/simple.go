package main

import (
	"os"

	"github.com/sg3des/ndia"
)

var (
	styleBox   = "fill:#212121; stroke:#000; stroke-width:2;"
	styleTitle = "text-anchor:middle; font-size:17px; fill:#fff; font-family: monospace;"
	styleLine  = "stroke:#25a; stroke-width:4;"
)

func main() {
	canvas := ndia.NewCanvas(500, 500)

	circle := ndia.NewCircle(250, 250, 50, styleBox)
	canvas.AddObject(circle)

	box0 := ndia.NewBox(250, 250, 50, 2, styleLine)
	canvas.AddObject(box0)

	text := ndia.NewText(250, 250, 7, "title", styleTitle)
	canvas.AddObject(text)

	canvas.Draw(os.Stdout)
}
