package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/juanjosegongi/automata-celular-musical/models"
	"golang.org/x/image/colornames"
)

func handleErr(err error) {
	log.Fatal(err)
}

// Main function for pixelgl
func run() {
	win, err := createWindow()
	if err != nil {
		handleErr(err)
	}

	universe := initializeUniverse()

	for !win.Closed() {
		draw(win, universe)
		update(universe)
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

func initializeUniverse() *models.Universe {
	universe := models.NewUniverse()

	for automatonCount := 0; automatonCount < 3; automatonCount++ {
		universe.AddAutomaton()
	}

	return universe
}

func draw(win *pixelgl.Window, universe *models.Universe) {
	win.Clear(colornames.Black)
	universe.Draw(win)
	win.Update()
}

func update(universe *models.Universe) {
	universe.Update()
}

func main() {
	pixelgl.Run(run)
}
