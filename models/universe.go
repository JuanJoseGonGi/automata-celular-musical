package models

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

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
func NewUniverse() *Universe {
	return &Universe{
		Automatons: generateInitialAutomatons(),
		Rules:      generateRules(),
	}
}

func generateInitialAutomatons() []*Automaton {
	automatons := []*Automaton{}

	for i := 0; i < 3; i++ {
		automatons = append(automatons, NewAutomaton())
	}

	return automatons
}

func splitRules(permutations [][]int) Rules {
	rand.Seed(time.Now().UnixNano())
	rules := Rules{}

	for _, permutation := range permutations {
		keySlc := []string{}

		for _, number := range permutation {
			keySlc = append(keySlc, strconv.Itoa(number))
		}

		key := strings.Join(keySlc, "")

		rules[key] = rand.Intn(6)
	}

	return rules
}

func possibleStates() []int {
	values := []int{}
	for i := 0; i < 5; i++ {
		values = append(values, i)
	}

	return values
}

func generatePermutation(permutationCount []int, permutation []int, values []int) []int {
	// generate permutaton
	for i, x := range permutationCount {
		permutation[i] = values[x]
	}
	return permutation
}

func generateRules() Rules {
	permutations := [][]int{}

	values := possibleStates()

	permutationCount := make([]int, len(values))
	permutation := make([]int, len(values))

	for {
		permutation = generatePermutation(permutationCount, permutation, values)
		permutations = append(permutations, permutation)
		// increment permutation number
		for i := 0; ; {
			permutationCount[i]++
			if permutationCount[i] < len(values) {
				break
			}
			permutationCount[i] = 0
			i++
			if i == len(values) {
				return splitRules(permutations)
			}
		}
	}
}

// AddAutomaton adds a new Automaton
func (universe *Universe) AddAutomaton() {
	automaton := NewAutomaton()
	universe.Automatons = append(universe.Automatons, automaton)
}

// Update updates each automaton
func (universe *Universe) Update() {
	for _, automaton := range universe.Automatons {
		automaton.Update(universe.Rules)
	}
}

// Draw draws each automaton
func (universe *Universe) Draw(win *pixelgl.Window) {
	for _, automaton := range universe.Automatons {
		automaton.Draw(win)
	}
}
