package grid

import (
	"fmt"
	"testing"
)

func checkPath(g Interface, start, end Coordinate, path []Coordinate) bool {
	prev := start
	for _, c := range path {
		d := ManhattanDistance(g, prev, c)
		if d != 1 {
			return false
		}
		prev = c
	}

	if prev != end {
		return false
	}

	return true
}

func printPath(path []Coordinate) string {
	var result string
	for _, c := range path {
		result += fmt.Sprint(c)
	}
	return result
}

func TestAStarNegMapRows(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{-1, 1}
	c := Coordinate{0, 0}

	AStarForGrid(g, c, c)
	t.Error("AStarForGrid on illegal grid should panic.")
}

func TestAStarNegMapCols(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{1, -1}
	c := Coordinate{0, 0}

	AStarForGrid(g, c, c)
	t.Error("AStarForGrid on illegal grid should panic.")
}

func TestAStarZeroMapRows(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{0, 1}
	c := Coordinate{0, 0}

	AStarForGrid(g, c, c)
	t.Error("AStarForGrid on illegal grid should panic.")
}

func TestAStarZeroMapCols(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{1, 0}
	c := Coordinate{0, 0}

	AStarForGrid(g, c, c)
	t.Error("AStarForGrid on illegal grid should panic.")
}

func TestAStarIllegalStart(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{1, 1}
	s := Coordinate{1, 0}
	e := Coordinate{0, 0}

	AStarForGrid(g, s, e)
	t.Error("AStarForGrid on illegal start should panic.")
}

func TestAStarIllegalEnd(t *testing.T) {
	defer func() { recover() }()
	g := &testGrid{1, 1}
	s := Coordinate{0, 0}
	e := Coordinate{0, 1}

	AStarForGrid(g, s, e)
	t.Error("AStarForGrid on illegal end should panic.")
}

func TestAStar(t *testing.T) {
	g := &testGrid{20, 20}
	s := Coordinate{6, 7}
	e := Coordinate{12, 13}

	AStarForGrid(g, s, e)
	p, found := AStarForGrid(g, s, e)
	if !found || len(p) != 12 || !checkPath(g, s, e, p) {
		t.Error("AStarForGrid returned an incorrect path.")
		t.Error(printPath(p))
	}
}

func TestAStarStartAtEnd(t *testing.T) {
	g := &testGrid{3, 3}
	c := Coordinate{1, 1}

	p, found := AStarForGrid(g, c, c)
	if !found || len(p) != 0 || !checkPath(g, c, c, p) {
		t.Error("AStarForGrid returned an incorrect path.")
		t.Error(printPath(p))
	}
}

func TestAStarWrapping(t *testing.T) {
	g := &testGrid{20, 20}
	s := Coordinate{4, 5}
	e := Coordinate{19, 18}

	p, found := AStarForGrid(g, s, e)
	if !found || len(p) != 12 || !checkPath(g, s, e, p) {
		t.Error("AStarForGrid returned an incorrect path.")
		t.Error(printPath(p))
	}
}

type testGridBlocking struct {
	rows, cols int
	passable   map[Coordinate]bool
}

func (g *testGridBlocking) NumRows() int {
	return g.rows
}

func (g *testGridBlocking) NumCols() int {
	return g.cols
}

func (g *testGridBlocking) IsPassable(c Coordinate) bool {
	p, ok := g.passable[c]
	return !ok || p
}

func TestAStarBlocked(t *testing.T) {
	g := &testGridBlocking{3, 3, make(map[Coordinate]bool)}
	s := Coordinate{1, 1}
	e := Coordinate{2, 2}

	g.passable[Coordinate{0, 1}] = false
	g.passable[Coordinate{1, 0}] = false
	g.passable[Coordinate{1, 2}] = false
	g.passable[Coordinate{2, 1}] = false

	p, found := AStarForGrid(g, s, e)
	if found {
		t.Error("AStarForGrid returned an incorrect path.")
		t.Error(printPath(p))
	}
}

func TestAStarEndImpassible(t *testing.T) {
	g := &testGridBlocking{3, 3, make(map[Coordinate]bool)}
	s := Coordinate{0, 0}
	e := Coordinate{1, 1}

	g.passable[e] = false

	p, found := AStarForGrid(g, s, e)
	if found {
		t.Error("AStarForGrid returned an incorrect path.")
		t.Error(printPath(p))
	}
}

func TestAStarWrappingAroundImpassible(t *testing.T) {
	g := &testGridBlocking{4, 4, make(map[Coordinate]bool)}
	s := Coordinate{1, 1}
	e := Coordinate{3, 3}

	g.passable[Coordinate{2, 0}] = false
	g.passable[Coordinate{2, 1}] = false
	g.passable[Coordinate{2, 3}] = false
	g.passable[Coordinate{0, 2}] = false
	g.passable[Coordinate{1, 2}] = false
	g.passable[Coordinate{3, 2}] = false

	p, found := AStarForGrid(g, s, e)
	if !found || len(p) != 4 || !checkPath(g, s, e, p) {
		t.Error("AStarForGrid returned an incorrect path.")
		t.Error(printPath(p))
	}
}
