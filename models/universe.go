package models

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/faiface/pixel/pixelgl"
)

// Rules handles a group of Rules
type Rules map[string]int

// Universe is a struct to store the application universe
type Universe struct {
	Automatons []*Automaton
	Rules      Rules
}

// NewUniverse creates a new universe
func NewUniverse(config Config) (*Universe, error) {
	initialAutomatons, err := generateInitialAutomatons(config)
	if err != nil {
		return nil, err
	}

	return &Universe{
		Automatons: initialAutomatons,
		Rules:      generateRules(config.NotesAmount),
	}, nil
}

func generateInitialAutomatons(config Config) ([]*Automaton, error) {
	automatons := []*Automaton{}

	for index, instrumentName := range config.Instruments {
		automaton, err := NewAutomaton(instrumentName, index, config)
		if err != nil {
			return nil, err
		}

		automatons = append(automatons, automaton)
	}

	return automatons, nil
}

func splitRules(permutations [][]int, notesAmount int) Rules {
	rules := Rules{}

	for _, permutation := range permutations {
		keySlc := []string{}

		for _, number := range permutation {
			keySlc = append(keySlc, strconv.Itoa(number))
		}

		key := strings.Join(keySlc, "")

		rules[key] = rand.Intn(notesAmount)
	}

	return rules
}

func possibleStates(notesAmount int) []int {
	values := []int{}
	for i := 0; i < notesAmount; i++ {
		values = append(values, i)
	}

	return values
}

func generatePermutation(permutationCount []int, values []int) []int {
	permutation := make([]int, 3)

	// generate permutaton
	for i, x := range permutationCount {
		permutation[i] = values[x]
	}
	return permutation
}

func generateRules(notesAmount int) Rules {
	permutations := [][]int{}

	values := possibleStates(notesAmount)

	permutationCount := make([]int, 3)
	for {
		permutation := generatePermutation(permutationCount, values)
		permutations = append(permutations, permutation)

		for i := 0; ; {
			permutationCount[i]++
			if permutationCount[i] < len(values) {
				break
			}
			permutationCount[i] = 0
			i++
			if i == 3 {
				return splitRules(permutations, notesAmount)
			}
		}
	}
}

// AddAutomaton adds a new Automaton
func (universe *Universe) AddAutomaton(instrumentName string, index int, config Config) error {
	automaton, err := NewAutomaton(instrumentName, index, config)
	if err != nil {
		return err
	}
	universe.Automatons = append(universe.Automatons, automaton)

	return nil
}

// Update updates each automaton
func (universe *Universe) Update(notesAmount int) error {
	for _, automaton := range universe.Automatons {
		err := automaton.Update(universe.Rules, notesAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

// Draw draws each automaton
func (universe *Universe) Draw(win *pixelgl.Window) {
	for _, automaton := range universe.Automatons {
		automaton.Draw(win)
	}
}
