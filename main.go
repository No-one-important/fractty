// fractty - Mandelbrot set explorer in your terminal
// SPDX-FileCopyrightText: 2022 Roland G. McIntosh
// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/encoding"
	"github.com/gdamore/tcell/v2"
)

var style = tcell.StyleDefault
var quit chan bool

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	encoding.Register()

	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	screen.EnableMouse()
	screen.Clear()

	quit = make(chan bool)
	go pollEvents(screen)

	screen.Show()

	go func() {
		for {
			drawScreen(screen)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	<-quit
	screen.Fini()
}

type viewport struct {
	x0, x1, y0, y1 float64
}

// Whether or not the given point is in the Mandelbrot set
func isConvergent(ca, cb float64) (bool, int) {
	var a, b float64 = 0, 0
	max := 64
	var i int
	for i = 0; i < max; i++ {
		as, bs := a*a, b*b
		if as+bs > 2 {
			return false, i
		}
		a, b = as-bs+ca, 2*a*b+cb
	}
	return true, i
}
