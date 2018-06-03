package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Image is simple box
type Image struct {
	Box
	Source image.Image
}

// NewImage make new *ui.Image
func NewImage(w, h int, srcImg image.Image) *Image {
	r := image.Rect(0, 0, w, h)
	b := Box{r, nil, nil, nil, []Item{}, Callbacks{}, nil}
	i := &Image{b, srcImg}
	i.Box.Super = i
	return i
}

// Reflesh updates internal *ebiten.Image
func (i *Image) Reflesh() {
	srcImg, _ := ebiten.NewImageFromImage(i.Source, ebiten.FilterDefault)
	w, h := i.Size()
	srcW, srcH := srcImg.Bounds().Dx(), srcImg.Bounds().Dy()
	scaleW, scaleH := float64(w)/float64(srcW), float64(h)/float64(srcH)
	i.Image, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleW, scaleH)
	i.Image.DrawImage(srcImg, op)
}

// String for fmt.Stringer interface
func (i *Image) String() string {
	p := fmt.Sprintf("%p", i)[7:11]
	return fmt.Sprintf("Image[%s]%s", p, i.Box.Rect)
}
