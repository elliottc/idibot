package grid

import (
	"math"
)

// Location combines (Row, Col) coordinate pairs for use as keys in maps (and in a 1d array).
type Location int

// Coordinate specifies a location on a grid.
type Coordinate struct{ row, col int }

// Gives the shortest distance between two coordinates on a wrapping line of total distance @total.
func StraightDistance(coord1, coord2, total int) int {
	if total < 0 {
		panic("mapTotal must be postive.")
	}

	// Get the distance in one direction.
	d := (coord1 - coord2) % total
	if d < 0 {
		d = -d
	}

	// Get the distance in the other direction.
	d2 := total - d

	if d < d2 {
		return d
	}
	return d2
}

// Gives the manhattan distance between two points on a wrapping map of the given dimensions.
func ManhattanDistance(row1, col1, row2, col2, maxRow, maxCol int) int {
	return StraightDistance(row1, row2, maxRow) + StraightDistance(col1, col2, maxCol)
}

// Gives the euclidean distance between two points on a wrapping map of the given dimensions.
func EuclideanDistance(row1, col1, row2, col2, maxRow, maxCol int) float64 {
	r := float64(StraightDistance(row1, row2, maxRow))
	c := float64(StraightDistance(col1, col2, maxCol))
	return math.Sqrt(r*r + c*c)
}
