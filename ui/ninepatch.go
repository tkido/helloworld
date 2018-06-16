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
	s := []image.Rectangle{}
	for k := 0; k < 9; k++ {
		var x0, y0, x1, y1 int
		i, j := k%3, k/3
		switch i {
		case 0:
			x0, x1 = out.Min.X, in.Min.X
		case 1:
			x0, x1 = in.Min.X, in.Max.X
		case 2:
			x0, x1 = in.Max.X, out.Max.X
		}
		switch j {
		case 0:
			y0, y1 = out.Min.Y, in.Min.Y
		case 1:
			y0, y1 = in.Min.Y, in.Max.Y
		case 2:
			y0, y1 = in.Max.Y, out.Max.Y
		}
		s = append(s, image.Rect(x0, y0, x1, y1))
	}
	return s
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
	for k := 0; k < 9; k++ {
		op := &ebiten.DrawImageOptions{}
		op.SourceRect = &(s[k])
		op.GeoM.Scale(float64(t[k].Dx())/float64(s[k].Dx()), float64(t[k].Dy())/float64(s[k].Dy()))
		op.GeoM.Translate(float64(t[k].Min.X), float64(t[k].Min.Y))
		n.Image.DrawImage(src, op)
	}
}

// String for fmt.Stringer interface
func (n *Ninepatch) String() string {
	p := fmt.Sprintf("%p", n)[7:11]
	return fmt.Sprintf("Image[%s]%s", p, n.Box.Rect)
}
