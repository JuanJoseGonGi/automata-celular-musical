package models

import (
	"math/rand"

	"github.com/faiface/pixel/pixelgl"
)

// Automaton is a struct to represent an Automaton
type Automaton struct {
	Cells []*Cell
	Step  int
}

// NewAutomaton creates a new Automaton
func NewAutomaton() *Automaton {
	return &Automaton{
		Cells: generateCells(),
		Step:  0,
	}
}

func generateCells() []*Cell {
	cells := []*Cell{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 13; j++ {
			cell := NewCell(float64(30+(i*50)), float64(35+(j*50)), rand.Intn(6))
			cells = append(cells, cell)
		}
	}

	return cells
}

// Update updates the state of each cell
func (automaton *Automaton) Update(rules Rules) {
	for index, cell := range automaton.Cells {
		cell.PrevState = cell.State

		prevCellState := "0"
		if index > 0 {
			prevCellState = automaton.Cells[index-1].SPrevState()
		}

		currentCellState := cell.SState()

		nextCellState := "0"
		if index < len(automaton.Cells)-1 {
			nextCellState = automaton.Cells[index+1].SState()
		}

		rule := prevCellState + currentCellState + nextCellState

		cell.State = rules[rule]
	}
}

// Draw draws each cell
func (automaton *Automaton) Draw(win *pixelgl.Window) {
	for _, cell := range automaton.Cells {
		cell.Draw(win)
	}
}
