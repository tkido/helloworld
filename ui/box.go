package ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Item is ebiten UI item
type Item interface {
	Draw(*ebiten.Image, image.Point, image.Rectangle) error
	// Move(p image.Point) error
	Add(Item) error
	HandleMouseEvent(MouseEvent) (handled bool, err error)
}

// Box is simple box
type Box struct {
	Rect             image.Rectangle
	Color            color.Color
	Image            *ebiten.Image
	DrawImageOptions *ebiten.DrawImageOptions
	Children         []Item
}

// NewBox make new Box
func NewBox(x, y, w, h int, c color.Color) *Box {
	r := image.Rect(x, y, x+w, y+h)
	img, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	img.Fill(c)
	// opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(float64(x), float64(y))
	b := &Box{r, c, img, nil, []Item{}}
	return b
}

// Add append child item to item
func (b *Box) Add(c Item) error {
	b.Children = append(b.Children, c)
	return nil
}

// Draw draw box
func (b *Box) Draw(screen *ebiten.Image, origin image.Point, clip image.Rectangle) error {
	rect := b.Rect.Add(origin)
	clip = clip.Intersect(rect)
	if clip.Empty() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	p := rect.Min
	// clipped part of image
	if clip != rect {
		d := clip.Min.Sub(rect.Min)
		op.SourceRect = &image.Rectangle{d, d.Add(clip.Size())}
		p = p.Add(d)
	}
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(b.Image, op)

	for _, c := range b.Children {
		c.Draw(screen, origin.Add(b.Rect.Min), clip)
	}
	return nil
}

// String for fmt.Stringer interface
func (b *Box) String() string {
	p := fmt.Sprintf("%p", b)[7:11]
	return fmt.Sprintf("Box[%s]%s%s", p, b.Rect, ColorCode(b.Color))
}

// Move sets position
// func (b *Box) Move(x, y int) error {
// 	s := b.Rect.Size()
// 	b.Rect = image.Rect(x, y, x+s.X, y+s.Y)
// 	img, _ := ebiten.NewImage(s.X, s.Y, ebiten.FilterDefault)
// 	img.Fill(b.Color)
// 	opts := &ebiten.DrawImageOptions{}
// 	opts.GeoM.Translate(float64(x), float64(y))
// 	b.DrawImageOptions = opts
// 	return nil
// }

// HandleMouseEvent handle mouse event
func (b *Box) HandleMouseEvent(e MouseEvent) (handled bool, err error) {
	// out of range
	if !e.Point.In(b.Rect) {
		return
	}
	// children first because they are in front of parent
	for i := len(b.Children) - 1; 0 <= i; i-- {
		child := b.Children[i]
		// children are evaluated in reverse order
		// because that was added later is more front
		ok, err := child.HandleMouseEvent(e)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	// handle myself
	switch e.Type {
	case MouseDown:
		// fmt.Printf("Box[%p]:%s\n", b, e)
		m.Downed = &Downed{b, e.Point}
	case MouseUp:
		if m.Downed != nil {
			if m.Downed.Item == b {
				if IsCloseAsClick(e.Point, (*m.Downed).Point) {
					fmt.Printf("%s %s\n", b, "Clicked!!")
				}
			}
		}
		m.Downed = nil
		// fmt.Printf("Box[%p]:%s\n", b, e)
	}
	return true, nil
}
