package perlin

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/atotto/clipboard"
)

func TestPerlin(t *testing.T) {
	buf := bytes.Buffer{}
	for x := 0.0; x < 256.0; x += 0.1 {
		s := fmt.Sprintf("%f\t%f\n", x, fractal(x))
		buf.WriteString(s)
	}
	err := clipboard.WriteAll(buf.String())
	if err != nil {
		t.Error(err)
	}
}
func TestFloor(t *testing.T) {
	cases := []struct {
		Put  float64
		Want int
	}{
		{0.0, 0},
		{0.1, 0},
		{0.9999, 0},
		{1.0, 1},
		{1.0001, 1},
		{1.9999, 1},
		{-0.0, 0},
		{-0.1, -1},
		{-0.9999, -1},
		{-1.0, -1},
		{-1.0001, -2},
		{-1.9999, -2},
	}
	for _, c := range cases {
		got := int(math.Floor(c.Put))
		want := c.Want
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
