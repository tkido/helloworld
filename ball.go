package main

import (
	"bitbucket.org/tkido/helloworld/vector"
	"github.com/hajimehoshi/ebiten"
)

// Ball is ball
type Ball struct {
	R     float64
	P, V  vector.Vector
	Image *ebiten.Image
}

// NewBall is NewBall
func NewBall(r float64, p, v vector.Vector, i *ebiten.Image) *Ball {
	return &Ball{r, p, v, i}
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
	return
}

// Draw draw
func (b *Ball) Draw(target *ebiten.Image) (err error) {
	w, h := b.Image.Size()
	scaleX, scaleY := 2.0*b.R/float64(w), 2.0*b.R/float64(h)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(b.P.X-b.R, b.P.Y-b.R)
	target.DrawImage(b.Image, opts)
	return
}

// func (b *Ball) collisioned(o *Ball) bool {
// 	d := math.Sqrt(math.Pow(b.X-o.X, 2) + math.Pow(b.Y-o.Y, 2))
// 	return d <= b.R+o.R
// }
