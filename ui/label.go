package ui

import (
	"fmt"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/text"
)

// Align is align
type Align int

// Align
const (
	Left   Align = 0
	Top    Align = 0
	Right  Align = 1
	Bottom Align = 1
	Center Align = 2
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
	Align
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
func NewLabel(w, h int, text string, font int, align Align, color, bgColor color.Color) *Label {
	b := NewBox(w, h, bgColor)
	l := &Label{*b, text, font, color, align}
	l.Sub = l
	return l
}

// Reflesh updates internal *ebiten.Image
func (l *Label) Reflesh() {
	l.Box.Reflesh()
	f := m.FontManager.Fonts[l.FontType]

	left := 0
	w, _ := l.Size()
	length := font.MeasureString(f.Face, l.Text).Ceil()

	switch l.Align {
	case Left:
	case Right:
		left = w - length
	case Center:
		left = (w - length) / 2
	}
	text.Draw(l.Image, l.Text, f.Face, left, f.Ascent, l.FontColor)
}

// String for fmt.Stringer interface
func (l *Label) String() string {
	p := fmt.Sprintf("%p", l)[7:11]
	return fmt.Sprintf("Label[%s]%s:%s", p, l.Rect, string([]rune(l.Text)[:4]))
}
