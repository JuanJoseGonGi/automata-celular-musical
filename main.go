package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/juanjosegongi/automata-celular-musical/models"
	"golang.org/x/image/colornames"
)

var (
	config models.Config = models.Config{
		Speed:       1,
		Instruments: []string{},
		NotesAmount: 5,
	}
)

const (
	sampleRate = 44100
	bufferSize = 4410
)

func handleErr(err error) {
	log.Fatal(err)
}

func loadFile() error {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := loadFile()
	if err != nil {
		handleErr(err)
	}
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

	universe, err := models.NewUniverse(config)
	if err != nil {
		handleErr(err)
	}

	win.Clear(colornames.Black)
	universe.Draw(win)

	for !win.Closed() {
		win.Clear(colornames.Black)
		err = universe.Update(config.NotesAmount)
		if err != nil {
			handleErr(err)
		}
		universe.Draw(win)

		win.Update()
		time.Sleep(time.Second / time.Duration(2*config.Speed))
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
