package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/juanjosegongi/automata-celular-musical/models"
	"golang.org/x/image/colornames"
)

const (
	sampleRate = 44100
	bufferSize = 4410
)

func handleErr(err error) {
	log.Fatal(err)
}

// Main function for pixelgl
func run() {
	rand.Seed(time.Now().UnixNano())
	err := speaker.Init(sampleRate, bufferSize)
	if err != nil {
		handleErr(err)
	}

	win, err := createWindow()
	if err != nil {
		handleErr(err)
	}

	instruments := []string{"piano", "saxophone", "tuba"}

	universe, err := models.NewUniverse(instruments)
	if err != nil {
		handleErr(err)
	}

	win.Clear(colornames.Black)
	universe.Draw(win)

	for !win.Closed() {
		win.Clear(colornames.Black)
		err = universe.Update()
		if err != nil {
			handleErr(err)
		}
		universe.Draw(win)

		win.Update()
		time.Sleep(time.Second / 2)
	}
}

func createWindow() (*pixelgl.Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "Musical Cellular Automaton",
		Bounds: pixel.R(0, 0, 1024, 700),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, err
	}

	return win, nil
}

func main() {
	pixelgl.Run(run)
}
