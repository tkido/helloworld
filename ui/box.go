package ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Item is ebiten UI item
type Item interface {
	Draw(screen *ebiten.Image, origin image.Point, clip image.Rectangle)
	Move(x, y int)
	Resize(w, h int)
	Size() (w, h int)
	Add(x, y int, item Item)
	HandleMouseEvent(ev MouseEvent, origin image.Point, clip image.Rectangle) (handled bool)
}

// Box is simple box
type Box struct {
	Rect             image.Rectangle
	Color            color.Color
	Image            *ebiten.Image
	DrawImageOptions *ebiten.DrawImageOptions
	Children         []Item
	CallbacksToMounseEvents
}

// NewBox make new Box
func NewBox(w, h int, c color.Color) *Box {
	r := image.Rect(0, 0, w, h)
	b := &Box{r, c, nil, nil, []Item{}, CallbacksToMounseEvents{}}
	b.Reflesh()
	return b
}

// Reflesh updates internal Image *ebiten.Image
func (b *Box) Reflesh() {
	w, h := b.Size()
	b.Image, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	b.Image.Fill(b.Color)
}

// Add append child item to item
func (b *Box) Add(x, y int, c Item) {
	c.Move(x, y)
	b.Children = append(b.Children, c)
}

// Move move item. (x, y) is relative position from parent.
func (b *Box) Move(x, y int) {
	b.Rect = b.Rect.Add(image.Point{x, y})
}

// Resize resize item
func (b *Box) Resize(w, h int) {
	x, y := b.Rect.Min.X, b.Rect.Min.Y
	b.Rect = image.Rect(x, y, x+w, y+h)
	b.Reflesh()
}

// Size get size of item
func (b *Box) Size() (w, h int) {
	s := b.Rect.Size()
	return s.X, s.Y
}

// Draw draw box
func (b *Box) Draw(screen *ebiten.Image, origin image.Point, clip image.Rectangle) {
	rect := b.Rect.Add(origin)
	clip = clip.Intersect(rect)
	if clip.Empty() {
		return
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
	return
}

// String for fmt.Stringer interface
func (b *Box) String() string {
	p := fmt.Sprintf("%p", b)[7:11]
	return fmt.Sprintf("Box[%s]%s%s", p, b.Rect, ColorCode(b.Color))
}

// HandleMouseEvent handle mouse event
func (b *Box) HandleMouseEvent(ev MouseEvent, origin image.Point, clip image.Rectangle) (handled bool) {
	rect := b.Rect.Add(origin)
	clip = clip.Intersect(rect)
	if clip.Empty() {
		return
	}
	// out of range
	if !ev.Point.In(clip) {
		return
	}
	// children first because they are in front of parent
	for i := len(b.Children) - 1; 0 <= i; i-- {
		child := b.Children[i]
		// children are evaluated in reverse order
		// because that was added later is more front
		ok := child.HandleMouseEvent(ev, origin.Add(b.Rect.Min), clip)
		if ok {
			return true
		}
	}
	// handle myself
	if callBack, ok := b.CallbacksToMounseEvents[ev.Type]; ok {
		callBack(b)
		return true
	}
	switch ev.Type {
	case MouseDown:
		// fmt.Printf("Box[%p]:%s\n", b, e)
		m.Downed = &Downed{b, ev.Point}
	case MouseUp:
		if m.Downed != nil {
			if m.Downed.Item == b {
				if IsCloseAsClick(ev.Point, (*m.Downed).Point) {
					w, h := b.Size()
					b.Resize(w+10, h+10)
					fmt.Printf("%s %s\n", b, "Clicked!!")
				}
			}
		}
		m.Downed = nil
		// fmt.Printf("Box[%p]:%s\n", b, e)
	}
	return true
}
