package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/text"
)

// Texter has internal text as string
type Texter interface {
	SetText(text string)
	GetText() (text string)
}

// Label is simple box
type Label struct {
	Box
	Text      string
	FontType  int
	FontColor color.Color
}

// SetText set internal text string
func (l *Label) SetText(s string) {
	l.Text = s
	l.Dirty()
}

// GetText get internal text string
func (l *Label) GetText() string {
	return l.Text
}

// NewLabel make new *ui.Label
func NewLabel(w, h int, text string, font int, color, bgColor color.Color) *Label {
	b := NewBox(w, h, bgColor)
	l := &Label{*b, text, font, color}
	l.Sub = l
	return l
}

// Reflesh updates internal *ebiten.Image
func (l *Label) Reflesh() {
	l.Box.Reflesh()
	f := m.FontManager.Fonts[l.FontType]
	text.Draw(l.Image, l.Text, f.Face, 0, f.Ascent, l.FontColor)
}

// String for fmt.Stringer interface
func (l *Label) String() string {
	p := fmt.Sprintf("%p", l)[7:11]
	return fmt.Sprintf("Label[%s]%s:%s", p, l.Rect, l.Text[:4])
}
