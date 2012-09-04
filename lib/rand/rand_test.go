package rand

import (
	"testing"
)

func setWeights(a []float64, weight float64) {
	for i := 0; i < len(a); i++ {
		a[i] = weight
	}
}

func TestRandSelectNil(t *testing.T) {
	defer func() { recover() }()
	RandSelect(nil)
	t.Error("RandSelect(nil) should panic.")
}

func TestRandSelectZero(t *testing.T) {
	defer func() { recover() }()
	a := make([]float64, 0)
	RandSelect(a)
	t.Error("RandSelect on empty input should panic.")
}

func TestRandSelectOneZeroWeight(t *testing.T) {
	a := make([]float64, 1)
	i := RandSelect(a)
	if i != 0 {
		t.Error("RandSelect on one element should select the element.")
	}
}

func TestRandSelectOne(t *testing.T) {
	a := make([]float64, 1)
	setWeights(a, 7)
	i := RandSelect(a)
	if i != 0 {
		t.Error("RandSelect on one element should select the element.")
	}
}

func TestRandSelectOneNegWeight(t *testing.T) {
	defer func() { recover() }()
	a := make([]float64, 1)
	setWeights(a, -4)
	RandSelect(a)
	t.Error("RandSelect on negative-valued input should panic.")
}

func TestRandSelectManyZeroWeight(t *testing.T) {
	a := make([]float64, 3)
	i := RandSelect(a)
	if i < 0 || i >= 3 {
		t.Error("RandSelect failed to select an element in the range.")
	}
}

func TestRandSelectMany(t *testing.T) {
	a := make([]float64, 3)
	setWeights(a, 7)
	i := RandSelect(a)
	if i < 0 || i >= 3 {
		t.Error("RandSelect failed to select an element in the range.")
	}
}

func TestRandSelectManySomeZeroWeight(t *testing.T) {
	a := make([]float64, 3)
	setWeights(a, 7)
	a[1] = 0
	i := RandSelect(a)
	if i < 0 || i >= 3 || i == 1 {
		t.Error("RandSelect failed to select a valid element in the range.")
	}
}

func TestRandSelectManyNegWeight(t *testing.T) {
	defer func() { recover() }()
	a := make([]float64, 3)
	a[1] = -4
	RandSelect(a)
	t.Error("RandSelect on negative-valued input should panic.")
}
