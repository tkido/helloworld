package ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Widget is ebiten UI item
type Widget struct {
	Children []Widget
}

// Draw draw UI to image
func (u *Widget) Draw(screen *ebiten.Image) error {
	for _, c := range u.Children {
		c.Draw(screen)
	}
	return nil
}

// func (u *Widget) HandleEvent() error {
// 	for _, c := range u.Children {
// 		c.Draw(screen)
// 	}
// 	return nil
// }

// NewBox make new Box
func NewBox(x, y, w, h int, c color.Color) *Box {
	r := image.Rect(x, y, x+w, y+h)
	img, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	img.Fill(c)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	b := &Box{r, c, img, opts}
	return b
}

// Box is simple box
type Box struct {
	Rect             image.Rectangle
	Color            color.Color
	Image            *ebiten.Image
	DrawImageOptions *ebiten.DrawImageOptions
}

// Draw draw box
func (b *Box) Draw(screen *ebiten.Image) error {
	screen.DrawImage(b.Image, b.DrawImageOptions)
	return nil
}

func (b *Box) String() string {
	return fmt.Sprintf("Box %#v", b)
}

// SetPosition set position
func (b *Box) SetPosition(x, y int) error {
	s := b.Rect.Size()
	b.Rect = image.Rect(x, y, x+s.X, y+s.Y)
	img, _ := ebiten.NewImage(s.X, s.Y, ebiten.FilterDefault)
	img.Fill(b.Color)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	b.DrawImageOptions = opts
	return nil
}

func (b *Box) HandleMouseEvent(e MouseEvent) (handled bool, err error) {
	switch e.Type {
	case MouseDown, MouseUp:
		fmt.Printf("Box[%p]:%s\n", b, e)
	}
	return true, nil
}
