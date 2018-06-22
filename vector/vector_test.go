package vector

import "testing"

func TestVector(t *testing.T) {
	v1 := Vector{1, 2}
	v2 := Vector{1, 2}
	if v1 != v2 {
		t.Errorf("got %v want %v", v1, v2)
	}
}
