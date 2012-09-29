package grid

import (
	"math"
)

// Coordinate specifies a location on a grid. 
type Coordinate struct{ Row, Col int }

// This specifies the interface a grid must implement to work with these library methods.
type Interface interface {
	NumRows() int
	NumCols() int
	IsPassable(c Coordinate) bool
}

// Gives the shortest distance between two coordinates on a wrapping line of total distance @total.
func straightDistance(coord1, coord2, total int) int {
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

// TODO: does Interface get passed by reference?

// Gives the manhattan distance between two points on a wrapping map of the given dimensions.
func ManhattanDistance(g Interface, c1, c2 Coordinate) int {
	return straightDistance(c1.Row, c2.Row, g.NumRows()) + straightDistance(c1.Col, c2.Col, g.NumCols())
}

// Gives the euclidean distance between two points on a wrapping map of the given dimensions.
func EuclideanDistance(g Interface, c1, c2 Coordinate) float64 {
	r := float64(straightDistance(c1.Row, c2.Row, g.NumRows()))
	c := float64(straightDistance(c1.Col, c2.Col, g.NumCols()))
	return math.Sqrt(r*r + c*c)
}
