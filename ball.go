package main

import (
	"math"

	"bitbucket.org/tkido/helloworld/quadtree"
	"bitbucket.org/tkido/helloworld/vector"
	"github.com/hajimehoshi/ebiten"
)

// Ball is ball
type Ball struct {
	R           float64
	P, V        vector.Vector
	Image       *ebiten.Image
	CellNum     int
	IsCollision bool
}

// NewBall is NewBall
func NewBall(r float64, p, v vector.Vector, i *ebiten.Image) *Ball {
	return &Ball{r, p, v, i, -1, false}
}

// Update is update
func (b *Ball) Update() (err error) {
	b.P = b.P.Add(b.V)
	if b.P.X-b.R < 0 || screenWidth <= b.P.X+b.R {
		b.V.X *= -1
	}
	if b.P.Y-b.R < 0 || screenHeight <= b.P.Y+b.R {
		b.V.Y *= -1
	}
	b.IsCollision = false
	return
}

// Draw draw
func (b *Ball) Draw(target *ebiten.Image) (err error) {
	w, h := b.Image.Size()
	scaleX, scaleY := 2.0*b.R/float64(w), 2.0*b.R/float64(h)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(b.P.X-b.R, b.P.Y-b.R)
	if b.IsCollision {
		opts.ColorM.RotateHue(math.Pi)
	}
	target.DrawImage(b.Image, opts)
	return
}

// CheckCollision is check collision
func (b *Ball) CheckCollision(o *Ball) bool {
	count++
	d := math.Sqrt(math.Pow(b.P.X-o.P.X, 2) + math.Pow(b.P.Y-o.P.Y, 2))
	return d <= b.R+o.R
}

// Check is check
func (b *Ball) Check(c quadtree.Collisioner) bool {
	count++
	if b.IsCollision == true {
		return false
	}
	switch o := c.(type) {
	case *Ball:
		if b == o {
			return false
		}
		d := math.Sqrt(math.Pow(b.P.X-o.P.X, 2) + math.Pow(b.P.Y-o.P.Y, 2))
		b.IsCollision = d <= b.R+o.R
		return true
	default:
		return false
	}
}

// GetCellNum retruns CellNum
func (b *Ball) GetCellNum() int {
	return b.CellNum
}

// SetCellNum sets CellNum
func (b *Ball) SetCellNum(c int) {
	b.CellNum = c
}
