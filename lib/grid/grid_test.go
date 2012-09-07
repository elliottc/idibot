package grid

import (
	"testing"
)

func TestStraightDistanceNegTotal(t *testing.T) {
	defer func() { recover() }()
	straightDistance(1, 2, -4)
	t.Error("straightDistance on negative total should panic.")
}

func TestStraightDistanceZeroTotal(t *testing.T) {
	defer func() { recover() }()
	straightDistance(2, 1, 0)
	t.Error("straightDistance on zero total should panic.")
}

func TestStraightDistance(t *testing.T) {
	d := straightDistance(1, 2, 5)
	if d != 1 {
		t.Error("straightDistance did not return the correct distance.")
	}
}

func TestStraightDistanceWrap(t *testing.T) {
	d := straightDistance(1, 4, 5)
	if d != 2 {
		t.Error("straightDistance did not return the correct wrapping distance.")
	}
}

func TestStraightDistanceNeg(t *testing.T) {
	d := straightDistance(45, -65, 10)
	if d != 0 {
		t.Error("straightDistance did not return the correct modulo distance.")
	}
}

func TestStraightDistanceNegNeg(t *testing.T) {
	d := straightDistance(-1, -15, 10)
	if d != 4 {
		t.Error("straightDistance did not return the correct modulo distance.")
	}
}

type testGrid struct {
	rows, cols int
}

func (g *testGrid) NumRows() int {
	return g.rows
}

func (g *testGrid) NumCols() int {
	return g.cols
}

func (g *testGrid) IsPassable(c Coordinate) bool {
	return true
}

func TestManhattanDistanceNegRowMax(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{-1, 1}
	c := Coordinate{0, 0}
	ManhattanDistance(g, c, c)
	t.Error("ManhattanDistance on negative rowMax should panic.")
}

func TestManhattanDistanceNegColMax(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{1, -1}
	c := Coordinate{0, 0}
	ManhattanDistance(g, c, c)
	t.Error("ManhattanDistance on negative colMax should panic.")
}

func TestManhattanDistance(t *testing.T) {
	g := &testGrid{6, 7}
	c := Coordinate{2, 3}
	d := ManhattanDistance(g, c, c)
	if d != 0 {
		t.Error("ManhattanDistance did not return the correct distance.")
	}
}

func TestManhattanDistanceMirror(t *testing.T) {
	g := &testGrid{10, 10}
	c1 := Coordinate{2, 3}
	c2 := Coordinate{-18, -17}
	d := ManhattanDistance(g, c1, c2)
	if d != 0 {
		t.Error("ManhattanDistance did not return the correct modulo distance.")
	}
}

func TestEuclideanDistance0Row(t *testing.T) {
	g := &testGrid{8, 9}
	c1 := Coordinate{3, 4}
	c2 := Coordinate{3, 6}
	expected := 2.0
	actual := EuclideanDistance(g, c1, c2)
	if actual != expected {
		t.Error("EuclideanDistance did not return the correct distance. Expected: ", expected, " actual: ", actual)
	}
}

func TestEuclideanDistance0Col(t *testing.T) {
	g := &testGrid{8, 9}
	c1 := Coordinate{3, 4}
	c2 := Coordinate{6, 4}
	expected := 3.0
	actual := EuclideanDistance(g, c1, c2)
	if actual != expected {
		t.Error("EuclideanDistance did not return the correct distance. Expected: ", expected, " actual: ", actual)
	}
}

func TestEuclideanDistance(t *testing.T) {
	g := &testGrid{9, 9}
	c1 := Coordinate{0, 0}
	c2 := Coordinate{3, 4}
	expected := 5.0
	actual := EuclideanDistance(g, c1, c2)
	if actual != expected {
		t.Error("EuclideanDistance did not return the correct distance. Expected: ", expected, " actual: ", actual)
	}
}

func TestEuclideanDistanceWrapping(t *testing.T) {
	g := &testGrid{21, 31}
	c1 := Coordinate{15, 2}
	c2 := Coordinate{2, 27}
	expected := 10.0
	actual := EuclideanDistance(g, c1, c2)
	if actual != expected {
		t.Error("EuclideanDistance did not return the correct distance. Expected: ", expected, " actual: ", actual)
	}
}
