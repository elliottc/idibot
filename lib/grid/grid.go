package grid

import (
	"math"
)

// Location combines (Row, Col) coordinate pairs for use as keys in maps (and in a 1d array).
// TODO: get rid of this? at least move it back into map - this concent doesnt belong in grid.
type Location int

// Coordinate specifies a location on a grid.
// TODO: make row, col private?
type Coordinate struct{ Row, Col int }

// This specifies the interface a grid must implement to work with these library methods.
type Interface interface {
	NumRows() int
	NumCols() int
	IsPassable(c Coordinate) bool
}

// ToLocation returns a Location given an (Row, Col) pair
func ToLocation(g Interface, c Coordinate) Location {
	for c.Row < 0 {
		c.Row += g.NumRows()
	}
	for c.Row >= g.NumRows() {
		c.Row -= g.NumRows()
	}
	for c.Col < 0 {
		c.Col += g.NumCols()
	}
	for c.Col >= g.NumCols() {
		c.Col -= g.NumCols()
	}

	return Location(c.Row*g.NumCols() + c.Col)
}

// ToCoordinate returns an (Row, Col) pair given a Location
func ToCoordinate(g Interface, loc Location) Coordinate {
	row := int(loc) / g.NumCols()
	col := int(loc) % g.NumCols()
	return Coordinate{row, col}
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

// TODO: can i do mod N/2?

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
