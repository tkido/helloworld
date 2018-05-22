package perlin

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"bitbucket.org/tkido/gostock/my"
)

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
func TestPerlin(t *testing.T) {
	buf := bytes.Buffer{}
	for x := 0.0; x < 10.01; x += 0.01 {
		s := fmt.Sprintf("%f\t%f\n", x, perlin(x))
		buf.WriteString(s)
	}
	my.WriteFileForCopyPaste("r.html", buf.String())
}

var gradient []float64

func init() {
	rand.Seed(time.Now().UnixNano())
	gradient = []float64{}
	for i := 0; i < 100; i++ {
		gradient = append(gradient, (rand.Float64()-0.5)*2)
	}
}

func c(x float64) float64 {
	if x < -1 || 1 < x {
		return 0
	}
	return 1 - 3*math.Pow(x, 2) + 2*math.Pow(math.Abs(x), 3)
}

// -1 <= x < 1
func wavelet(i int, x float64) float64 {
	return c(x) * gradient[int(i)] * x
}

// 0 <= x < 1
func perlin(x float64) float64 {
	if x <= 0 {
		return 0
	}
	f := math.Floor(x)
	x = x - f
	i := int(f)
	return wavelet(i, x) + x*(wavelet(i+1, x-1)-wavelet(i, x))
}
