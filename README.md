# ndia - SVG diagram

## install

```
go get github.com/sg3des/ndia
```

## example

	canvas := ndia.NewCanvas(500, 500)

	circle := ndia.NewCircle(250, 250, 50, styleBox)
	canvas.AddObject(circle)

	box0 := ndia.NewBox(250, 250, 50, 2, styleLine)
	canvas.AddObject(box0)

	text := ndia.NewText(250, 250, 7, "title", styleTitle)
	canvas.AddObject(text)

	canvas.Draw(os.Stdout)