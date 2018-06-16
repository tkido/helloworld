package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Ninepatch is nine-patch image
type Ninepatch struct {
	Box
	Source      image.Image
	ContentArea image.Rectangle
}

// NewNinepatch make new *ui.Ninepatch
func NewNinepatch(w, h int, srcImg image.Image) *Ninepatch {
	b := NewBox(w, h, nil)
	n := &Ninepatch{*b, srcImg, image.Rect(8, 8, 24, 24)}
	n.Sub = n
	return n
}

func makeNineRects(out, in image.Rectangle) []image.Rectangle {
	rs := make([]image.Rectangle, 0, 9)
	ys := [4]int{out.Min.Y, in.Min.Y, in.Max.Y, out.Max.Y}
	xs := [4]int{out.Min.X, in.Min.X, in.Max.X, out.Max.X}
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			rs = append(rs, image.Rect(xs[x], ys[y], xs[x+1], ys[y+1]))
		}
	}
	return rs
}

// Reflesh updates internal *ebiten.Image
func (n *Ninepatch) Reflesh() {
	src, _ := ebiten.NewImageFromImage(n.Source, ebiten.FilterDefault)
	w, h := n.Size()
	n.Image, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	so, si := src.Bounds(), n.ContentArea
	s := makeNineRects(so, si)
	to := n.Image.Bounds()
	tiSize := image.Point{to.Dx() - s[0].Dx() - s[2].Dx(), to.Dy() - s[0].Dy() - s[6].Dy()}
	ti := image.Rectangle{si.Min, si.Min.Add(tiSize)}
	t := makeNineRects(to, ti)
	for i := 0; i < 9; i++ {
		op := &ebiten.DrawImageOptions{}
		op.SourceRect = &(s[i])
		op.GeoM.Scale(float64(t[i].Dx())/float64(s[i].Dx()), float64(t[i].Dy())/float64(s[i].Dy()))
		op.GeoM.Translate(float64(t[i].Min.X), float64(t[i].Min.Y))
		n.Image.DrawImage(src, op)
	}
}

// String for fmt.Stringer interface
func (n *Ninepatch) String() string {
	p := fmt.Sprintf("%p", n)[7:11]
	return fmt.Sprintf("Image[%s]%s", p, n.Box.Rect)
}
