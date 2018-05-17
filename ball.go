package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

// Drawable is drawable object
type Drawable interface {
	Draw() error
}

// Ball is ball
type Ball struct {
	X, Y, R float64
	Image   *ebiten.Image
}

// NewBall is NewBall
func NewBall(x, y, r float64, i *ebiten.Image) *Ball {
	return &Ball{x, y, r, i}
}

// Draw draw
func (b *Ball) Draw(target *ebiten.Image) (err error) {
	w, h := b.Image.Size()
	scaleX, scaleY := 2.0*b.R/float64(w), 2.0*b.R/float64(h)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(b.X-b.R, b.Y-b.R)
	target.DrawImage(b.Image, opts)
	return
}

func (b *Ball) collisioned(o *Ball) bool {
	d := math.Sqrt(math.Pow(b.X-o.X, 2) + math.Pow(b.Y-o.Y, 2))
	return d <= b.R+o.R
}
