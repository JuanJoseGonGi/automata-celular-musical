package models

import (
	"math/rand"

	"github.com/faiface/pixel/pixelgl"
)

// Automaton is a struct to represent an Automaton
type Automaton struct {
	Cells      [][]*Cell
	CellIndex  int
	Step       int
	Index      int
	Instrument Instrument
}

// NewAutomaton creates a new Automaton
func NewAutomaton(instrumentName string, index int, speed float64) (*Automaton, error) {
	instrument, err := NewInstrument(instrumentName, speed)
	if err != nil {
		return nil, err
	}

	return &Automaton{
		Cells:      generateCells(index),
		CellIndex:  0,
		Step:       0,
		Index:      index,
		Instrument: instrument,
	}, nil
}

func generateCells(index int) [][]*Cell {
	cells := [][]*Cell{}

	for i := 0; i < 13; i++ {
		row := []*Cell{}

		for j := 0; j < 5; j++ {
			cell := NewCell(float64(index*280+110+(j*50)), float64(35+(i*50)), 0)
			row = append(row, cell)
		}
		cells = append(cells, row)
	}

	return cells
}

func (automaton *Automaton) playAndIncreaseCellIndex() error {
	step := automaton.Step % 13
	row := automaton.Cells[step]

	err := automaton.Instrument.playNote(row[automaton.CellIndex].State)
	if err != nil {
		return err
	}

	automaton.CellIndex++
	if automaton.CellIndex == 5 {
		automaton.Step++
		automaton.CellIndex = 0
	}

	return nil
}

// Update updates the state of each cell
func (automaton *Automaton) Update(rules Rules) error {
	step := automaton.Step % 13
	row := automaton.Cells[step]

	if automaton.Step == 0 {
		row[automaton.CellIndex].State = rand.Intn(5)
		return automaton.playAndIncreaseCellIndex()
	}

	var prevRow []*Cell
	if step == 0 {
		prevRow = automaton.Cells[len(automaton.Cells)-1]
	} else {
		prevRow = automaton.Cells[step-1]
	}

	prevCellState := "0"
	if automaton.CellIndex > 0 {
		prevCellState = prevRow[automaton.CellIndex-1].SState()
	}

	currentCellState := prevRow[automaton.CellIndex].SState()

	nextCellState := "0"
	if automaton.CellIndex < len(prevRow)-1 {
		nextCellState = prevRow[automaton.CellIndex+1].SState()
	}

	rule := prevCellState + currentCellState + nextCellState
	row[automaton.CellIndex].State = rules[rule]

	return automaton.playAndIncreaseCellIndex()
}

// Draw draws each cell
func (automaton *Automaton) Draw(win *pixelgl.Window) {
	for _, row := range automaton.Cells {
		for _, cell := range row {
			cell.Draw(win, row)
		}
	}
}
