// +build ignore

package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/eaburns/gui/thread0"
	"github.com/eaburns/gui/ui"
	"github.com/eaburns/gui/ui/sdl"
)

const (
	width   = 640
	height  = 480
	imgPath = "ui/gopher.png"
)

func main() {
	if err := sdl.Init(20); err != nil {
		panic(err)
	}
	go mainFunc()
	thread0.Hijack()
}

func mainFunc() {
	win, err := sdl.NewWindow("Test", width, height)
	if err != nil {
		panic(err)
	}

	canvas := ui.NewCanvas()
	img, err := ui.LoadPng(imgPath)
	if err != nil {
		panic(err)
	}
	img.Width = 100
	img.Height = 100

	tick := time.NewTicker(20 * time.Millisecond)
	for {
		select {
		case ev, ok := <-win.Events():
			if !ok {
				os.Exit(0)
			}
			fmt.Printf("%T%v\n", ev, ev)
			if _, ok := ev.(ui.WinClose); ok {
				win.Close()
			}
		case <-tick.C:
			canvas.Clear(color.Black)
			canvas.FillRect(10, 10, 20, 50, color.RGBA{R: 255, A: 255})
			canvas.FillRect(100, 100, 50, 50, color.RGBA{B: 255, A: 255})
			canvas.FillRect(200, 200, 100, 100, color.RGBA{G: 255, A: 255})

			canvas.StrokeLine(color.RGBA{R: 255, A: 255}, 1,
				[2]float32{width - 1, 0},
				[2]float32{0, height - 1})

			canvas.DrawImage(200, 200, img)
			win.Present()
		}
	}
	panic("Unreachable")
}
