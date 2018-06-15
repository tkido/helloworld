package ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Box is simple box
type Box struct {
	Rect             image.Rectangle
	Color            color.Color
	Image            *ebiten.Image
	DrawImageOptions *ebiten.DrawImageOptions
	Parent           Item
	Children         []Item
	Sub              Item
	dirty            bool
	Callbacks
}

// NewBox make new Box
func NewBox(w, h int, c color.Color) *Box {
	b := &Box{
		Rect:             image.Rect(0, 0, w, h),
		Color:            c,
		Image:            nil,
		DrawImageOptions: nil,
		Parent:           nil,
		Children:         []Item{},
		Callbacks:        Callbacks{},
		Sub:              nil,
		dirty:            true,
	}
	b.Sub = b
	return b
}

// Reflesh updates internal *ebiten.Image
func (b *Box) Reflesh() {
	w, h := b.Size()
	b.Image, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	if b.Color != nil && b.Color != color.Transparent {
		b.Image.Fill(b.Color)
	}
}

// Dirty set item dirty
func (b *Box) Dirty() {
	if b.dirty == true {
		return
	}
	b.dirty = true
	if b.Parent != nil {
		b.Parent.Dirty()
	}
}

// Add append child item to item
func (b *Box) Add(x, y int, c Item) {
	c.SetParent(b.Sub)
	b.Children = append(b.Children, c)
	c.Move(x, y)
}

// SetParent set parent
func (b *Box) SetParent(i Item) {
	b.Parent = i
}

// Move move item. (x, y) is relative position from parent.
func (b *Box) Move(x, y int) {
	w, h := b.Size()
	b.Rect = image.Rect(x, y, x+w, y+h)
	if b.Parent != nil {
		b.Parent.Dirty()
	}
}

// Position return relative position from parent Item
func (b *Box) Position() (x, y int) {
	min := b.Rect.Min
	return min.X, min.Y
}

// Resize resize item
func (b *Box) Resize(w, h int) {
	x, y := b.Rect.Min.X, b.Rect.Min.Y
	b.Rect = image.Rect(x, y, x+w, y+h)
	b.Dirty()
}

// Size get size of item
func (b *Box) Size() (w, h int) {
	s := b.Rect.Size()
	return s.X, s.Y
}

func copyDrawImageOptions(op ebiten.DrawImageOptions) ebiten.DrawImageOptions {
	if op.SourceRect != nil {
		// log.Printf("%p", op.SourceRect)
		sourceRect := *op.SourceRect
		op.SourceRect = &sourceRect
		// log.Printf("%p", op.SourceRect)
		// log.Printf("ok?")
	}
	return op
}

// Draw draw box
func (b *Box) Draw(target *ebiten.Image) {
	if b.dirty {
		b.Sub.Reflesh()
		b.dirty = false
		for _, c := range b.Children {
			c.Draw(b.Image)
		}
	}
	op := &ebiten.DrawImageOptions{}
	if o := b.DrawImageOptions; o != nil {
		copied := copyDrawImageOptions(*o)
		op = &copied
	}
	x, y := b.Position()
	op.GeoM.Translate(float64(x), float64(y))
	target.DrawImage(b.Image, op)
}

// String for fmt.Stringer
func (b *Box) String() string {
	p := fmt.Sprintf("%p", b)[7:11]
	return fmt.Sprintf("Box[%s]%s%s", p, b.Rect, ColorCode(b.Color))
}

// Call callback function if it exists
func (b *Box) Call(t EventType) {
	// log.Printf("%s %s", b.Sub, t)
	if c, ok := b.Callbacks[t]; ok {
		c(b.Sub)
	}
}

func (b *Box) mouseOn() {
	if m.OnItem != nil && m.OnItem != b.Sub {
		m.OnItem.Call(MouseOut)
	}
	b.Call(MouseOn)
	if m.OnItem != b.Sub {
		b.Call(MouseOver)
	}
	m.OnItem = b.Sub
	b.mouseIn()
}

func (b *Box) mouseIn() {
	b.Call(MouseIn)
	if _, ok := m.InItems[b.Sub]; !ok {
		b.Call(MouseEnter)
	}
	m.InItems[b.Sub] = m.Now
}

// HandleMouseEvent handle mouse event
func (b *Box) HandleMouseEvent(ev MouseEvent, origin image.Point, clip image.Rectangle) (handled bool) {
	rect := b.Rect.Add(origin)
	clip = clip.Intersect(rect)
	if !ev.Point.In(clip) {
		return
	}
	// at first, children are evaluated in reverse order. because child added later is at more front
	for i := len(b.Children) - 1; 0 <= i; i-- {
		if handled := b.Children[i].HandleMouseEvent(ev, origin.Add(b.Rect.Min), clip); handled {
			b.mouseIn()
			return true
		}
	}
	// handle by myself
	b.mouseOn()
	for i := 0; i < 3; i++ {
		down, up, click, doubleClick := LeftDown+EventType(i), LeftUp+EventType(i), LeftClick+EventType(i), LeftDoubleClick+EventType(i)
		switch ev.Moves[i] {
		case Down:
			b.Call(down)
			m.Downed[i] = &MouseRecord{b, ev.Point, m.Now}
		case Up:
			b.Call(up)
			// isClick?
			if m.Downed[i] != nil && m.Downed[i].Item == b && isCloseEnough(ev.Point, m.Downed[i].Point) {
				if m.Clicked[i] != nil {
					// isDoubleClick?
					if m.Clicked[i].Item == b && m.Now-m.Clicked[i].Frame <= m.DoubleClickInterval && isCloseEnough(ev.Point, m.Clicked[i].Point) {
						b.Call(doubleClick)
						m.Clicked[i] = nil
					} else {
						m.Clicked[i].Item.Call(click)
						m.Clicked[i] = &MouseRecord{b, ev.Point, m.Now}
					}
				} else if _, ok := b.Callbacks[doubleClick]; ok {
					m.Clicked[i] = &MouseRecord{b, ev.Point, m.Now}
				} else {
					b.Call(click)
				}
			}
			m.Downed[i] = nil
		}
	}
	return true
}
