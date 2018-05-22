package perlin

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/atotto/clipboard"
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
	for x := 0.0; x < 256.0; x += 0.1 {
		s := fmt.Sprintf("%f\t%f\n", x, fractal(x))
		buf.WriteString(s)
	}
	err := clipboard.WriteAll(buf.String())
	if err != nil {
		t.Error(err)
	}
}

var maxLevel = 6
var gradients [][]float64

func init() {
	rand.Seed(time.Now().UnixNano())
	gradients = [][]float64{}
	for lv := 0; lv <= maxLevel; lv++ {
		gradient := []float64{}
		for i := 0; i <= 256; i++ {
			gradient = append(gradient, (rand.Float64()-0.5)*2)
		}
		gradients = append(gradients, gradient)
	}

}

func c(x float64) float64 {
	return 1 - 6*math.Abs(math.Pow(x, 5)) + 15*math.Pow(x, 4) - 10*math.Abs(math.Pow(x, 3))
}

// -1 <= x < 1
func w(lv, i int, x float64) float64 {
	return c(x) * gradients[lv][int(i)] * x
}

// 0 <= x < 1
func perlin(lv int, x float64) float64 {
	if x < 0 {
		return 0
	}
	f := math.Floor(x)
	x = x - f
	i := int(f)
	return w(lv, i, x) + x*(w(lv, i+1, x-1)-w(lv, i, x))
}

func fractal(x float64) float64 {
	var scale float64
	var rst float64
	for lv := 0; lv <= maxLevel; lv++ {
		scale = float64(int(1 << uint(lv)))
		rst += perlin(lv, x/scale) * scale
	}
	return rst
}
