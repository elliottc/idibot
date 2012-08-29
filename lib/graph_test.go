package lib

import (
	"testing"
)

func TestStraightDistanceNegTotal(t *testing.T) {
	defer func() { recover() }()
	StraightDistance(1, 2, -4)
	t.Error("StraightDistance on negative total should panic.")
}

func TestStraightDistanceZeroTotal(t *testing.T) {
	defer func() { recover() }()
	StraightDistance(2, 1, 0)
	t.Error("StraightDistance on zero total should panic.")
}

func TestStraightDistance(t *testing.T) {
	d := StraightDistance(1, 2, 5)
	if d != 1 {
		t.Error("StraightDistance did not return the correct distance.")
	}
}

func TestStraightDistanceWrap(t *testing.T) {
	d := StraightDistance(1, 4, 5)
	if d != 2 {
		t.Error("StraightDistance did not return the correct wrapping distance.")
	}
}

func TestStraightDistanceNeg(t *testing.T) {
	d := StraightDistance(45, -65, 10)
	if d != 0 {
		t.Error("StraightDistance did not return the correct modulo distance.")
	}
}

func TestStraightDistanceNegNeg(t *testing.T) {
	d := StraightDistance(-1, -15, 10)
	if d != 4 {
		t.Error("StraightDistance did not return the correct modulo distance.")
	}
}

func TestManhattanDistanceNegRowMax(t *testing.T) {
	defer func() { recover() }()
	ManhattanDistance(0, 0, 0, 0, -1, 1)
	t.Error("ManhattanDistance on negative rowMax should panic.")
}

func TestManhattanDistanceNegColMax(t *testing.T) {
	defer func() { recover() }()
	ManhattanDistance(0, 0, 0, 0, 1, -1)
	t.Error("ManhattanDistance on negative colMax should panic.")
}

func TestManhattanDistance(t *testing.T) {
	d := ManhattanDistance(2, 3, 2, 3, 6, 7)
	if d != 0 {
		t.Error("ManhattanDistance did not return the correct distance.")
	}
}

func TestManhattanDistanceMirror(t *testing.T) {
	d := ManhattanDistance(2, 3, -18, -17, 10, 10)
	if d != 0 {
		t.Error("ManhattanDistance did not return the correct modulo distance.")
	}
}

// TODO: euclidean distance
