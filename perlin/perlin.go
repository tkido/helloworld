package perlin

import (
	"math"
	"math/rand"
	"time"
)

const (
	max      = 256
	maxLevel = 6
)

var gradients [][]float64

func init() {
	rand.Seed(time.Now().UnixNano())
	gradients = [][]float64{}
	for lv := 0; lv <= maxLevel; lv++ {
		scale := 1 << uint(lv)
		gradient := []float64{}
		for i := 0; i <= max/scale; i++ {
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
	return c(x) * gradients[lv][i] * x
}

func perlin(lv int, x float64) float64 {
	if x < 0 {
		return 0
	}
	f := math.Floor(x)
	x = x - f
	i := int(f)
	w0 := w(lv, i, x)
	w1 := w(lv, i+1, x-1)
	return w0 + x*(w1-w0)
}

func fractal(x float64) float64 {
	var rst float64
	for lv := 0; lv <= maxLevel; lv++ {
		scale := math.Exp2(float64(lv))
		rst += perlin(lv, x/scale) * scale
	}
	return rst
}
