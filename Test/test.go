package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var p float64 = 3.14159
var cc float64 = 3.2

func ring(d float64) string {
	v := (d * p) / cc
	return fmt.Sprint("Количество циклов: ", v, " Уголо поворота стола: ", 360/(v*2))
}

func loop(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			//Если убрать эту строчку, то окно будет открываться снова и снова.
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			l := material.H1(th, ring(120))
			maroon := color.NRGBA{R: 0, G: 200, B: 0, A: 255}
			l.Color = maroon
			l.Alignment = text.Middle
			l.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
