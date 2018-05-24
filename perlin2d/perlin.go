package perlin2d

import (
	"math"
	"math/rand"
	"time"
)

const (
	max      = 32
	maxLevel = 5
)

// Grad is gradient
type Grad struct {
	X, Y float64
}

var gradients [][][]Grad

func init() {
	rand.Seed(time.Now().UnixNano())
	gradients = [][][]Grad{}
	for lv := 0; lv <= maxLevel; lv++ {
		scale := 1 << uint(lv)
		gsY := [][]Grad{}
		for j := 0; j <= max/scale; j++ {
			gsX := []Grad{}
			for i := 0; i <= max/scale; i++ {
				gsX = append(gsX, Grad{randGrad(), randGrad()})
			}
			gsY = append(gsY, gsX)
		}
		gradients = append(gradients, gsY)
	}
}

func randGrad() float64 {
	return (rand.Float64() - 0.5) * 2
}

func c(x, y float64) float64 {
	return c1d(x) * c1d(y)
}
func c1d(x float64) float64 {
	return 1 - 6*math.Abs(math.Pow(x, 5)) + 15*math.Pow(x, 4) - 10*math.Abs(math.Pow(x, 3))
}

// -1 <= x < 1
func w(lv, i, j int, x, y float64) float64 {
	grad := gradients[lv][j][i]
	return c(x, y) * (grad.X*x + grad.Y*y)
}

func perlin(lv int, x, y float64) float64 {
	if x < 0 || y < 0 {
		return 0
	}
	fx, fy := math.Floor(x), math.Floor(y)
	x, y = x-fx, y-fy
	i, j := int(fx), int(fy)
	w00 := w(lv, i, j, x, y)
	w01 := w(lv, i+1, j, x-1, y)
	w10 := w(lv, i, j+1, x, y-1)
	w11 := w(lv, i+1, j+1, x-1, y-1)

	w0x := w00 + x*(w01-w00)
	w1x := w10 + x*(w11-w10)
	return w0x + y*(w1x-w0x)
}

func fractal(x, y float64) float64 {
	var rst float64
	for lv := 0; lv <= maxLevel; lv++ {
		scale := math.Exp2(float64(lv))
		rst += perlin(lv, x/scale, y/scale) * scale
	}
	return rst
}

// Fractal is fractal
func Fractal(x, y float64) float64 {
	return fractal(x, y)
}
