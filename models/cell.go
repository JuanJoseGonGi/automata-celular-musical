package models

import (
	"image/color"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	cellSize = 30
)

// Cell is a struct to handle a cell
type Cell struct {
	PosX      float64
	PosY      float64
	Size      float64
	State     int
	PrevState int
}

// NewCell creates a new cell
func NewCell(posX float64, posY float64, state int) *Cell {
	return &Cell{
		PosX:  posX,
		PosY:  posY,
		Size:  cellSize,
		State: state,
	}
}

// SState returns the current state as string
func (cell *Cell) SState() string {
	return strconv.Itoa(cell.State)
}

// Draw draws the cell on win
func (cell *Cell) Draw(win *pixelgl.Window, row []*Cell) {
	imd := imdraw.New(nil)

	imd.SetColorMask(cell.getColorFromState())
	imd.Push(pixel.V(cell.PosX, cell.PosY), pixel.V(cell.PosX+cell.Size, cell.PosY+cell.Size))
	imd.Rectangle(0)

	imd.Draw(win)
}

func (cell *Cell) getColorFromState() color.Color {
	switch cell.State {
	case 1:
		return colornames.Blue
	case 2:
		return colornames.Red
	case 3:
		return colornames.Green
	case 4:
		return colornames.Yellow
	case 5:
		return colornames.Orange
	}

	return color.White
}
