package idibot

import (
	"math/rand"
)

// given a list of weights, randomly choose the index of one
func RandSelect(weights []float64) int {
	var total float64
	for _, w := range weights {
		if w < 0 {
			panic("Cannot select a negative weight.")
		}
		total += w
	}

	var acc float64
	r := rand.Float64() * total
	for i, w := range weights {
		acc += w
		if r <= acc {
			return i
		}
	}

	panic("Failed to select an index.")
}
