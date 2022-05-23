// fractty - Mandelbrot set explorer in your terminal
// SPDX-FileCopyrightText: 2022 Roland G. McIntosh
// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"fmt"
	"os"
	"time"
	"flag"

	"github.com/gdamore/tcell/encoding"
	"github.com/gdamore/tcell/v2"
)

var style = tcell.StyleDefault
var quit chan bool
var fractal *string

func main() {
	fractal = flag.String("f", "mandelbrot", "Fractal type")
	flag.Parse()

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

func isConvergent(ca, cb float64) (bool, int) {
	switch *fractal {
		case "julia":
			return isConvergentJulia(ca, cb)
		default:
			return isConvergentMandelbrot(ca, cb)
	}
}

func isConvergentJulia(ca, cb float64) (bool, int) {
	maxIterations := 64
	za := 0.0
	zb := 0.0
	for i := 0; i < maxIterations; i++ {
		za = za*za - zb*zb + ca
		zb = 2*za*zb + cb
		if za*za+zb*zb > 4 {
			return false, i
		}
	}
	return true, 0
}

func isConvergentMandelbrot(ca, cb float64) (bool, int) {
	maxIterations := 64
	z := complex(ca, cb)
	c := complex(ca, cb)
	for i := 0; i < maxIterations; i++ {
		z = z*z + c
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			return false, i
		}
	}
	return true, maxIterations
}