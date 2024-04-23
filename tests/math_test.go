package math_test

import (
	"testing"

	"math"
)

func TestAbs(t *testing.T) {
	a := -1.0
	b := math.Abs(a)
	if b != 1 {
		t.Errorf("Abs(%f) = %f; want 1", a, b)
	}
}
