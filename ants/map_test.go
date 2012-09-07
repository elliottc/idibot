package ants

import (
	"idibot/lib/grid"
	"testing"
)

func TestMap(t *testing.T) {
	m := NewMap(4, 3)
	m.Reset()
	if m.String() != `. . . 
. . . 
. . . 
. . . 
` {
		t.Errorf("map is wrong size, got `%s`", m)
	}

	loc := grid.ToLocation(m, grid.Coordinate{3, 2})
	c := grid.ToCoordinate(m, loc)
	if c.Row != 3 || c.Col != 2 {
		t.Errorf("conversion broken, got (%v, %v), wanted (3, 2)", c.Row, c.Col)
	}

	loc2 := grid.ToLocation(m, grid.Coordinate{3, -1})
	if loc2 != loc {
		t.Errorf("from xy broken, got (%v), wanted (%v)", loc2, loc)
	}

	n := grid.ToLocation(m, grid.Coordinate{2, 2})
	s := grid.ToLocation(m, grid.Coordinate{4, 2})
	e := grid.ToLocation(m, grid.Coordinate{3, 3})
	w := grid.ToLocation(m, grid.Coordinate{3, 1})

	if n != m.Move(loc, North) {
		t.Errorf("Move north is broken")
	}
	if s != m.Move(loc, South) {
		t.Errorf("Move south is broken")
	}
	if e != m.Move(loc, East) {
		t.Errorf("Move east is broken")
	}
	if w != m.Move(loc, West) {
		t.Errorf("Move west is broken")
	}

	m.AddAnt(n, MY_ANT)
	m.AddAnt(s, ANT_1)
	m.AddAnt(w, MY_HILL)
	m.AddAnt(e, MY_OCCUPIED_HILL)
	m.AddHill(grid.ToLocation(m, grid.Coordinate{0, 0}), HILL_1)
	m.AddHill(grid.ToLocation(m, grid.Coordinate{1, 0}), HILL_1)
	m.AddAnt(grid.ToLocation(m, grid.Coordinate{1, 0}), ANT_1)

	if m.String() != `1 . b 
B . . 
. . a 
A 0 . 
` {
		t.Errorf("map put ants in wrong place, got `%s`", m)
	}
}
