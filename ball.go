package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

// Drawable is drawable object
type Drawable interface {
	Draw() error
}

// DrawableData is DrawableData
type DrawableData struct {
	Target *ebiten.Image
	Source *ebiten.Image
}

// Ball is ball
type Ball struct {
	X, Y, R float64
	DrawableData
}

// NewBall is NewBall
func NewBall(x, y, r float64, d DrawableData) *Ball {
	return &Ball{x, y, r, d}
}

// Draw draw
func (b *Ball) Draw() (err error) {
	w, h := b.Source.Size()
	scaleX, scaleY := 2.0*b.R/float64(w), 2.0*b.R/float64(h)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(b.X-b.R, b.Y-b.R)
	err = b.Target.DrawImage(b.Source, opts)
	if err != nil {
		return
	}
	return
}

func (b *Ball) collisioned(o *Ball) bool {
	d := math.Sqrt(math.Pow(b.X-o.X, 2) + math.Pow(b.Y-o.Y, 2))
	return d <= b.R+o.R
}
