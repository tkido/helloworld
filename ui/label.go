package ui

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/text"
)

// Label is simple box
type Label struct {
	Box
	Text      string
	Face      font.Face
	FontColor color.Color
	FontSize  int
}

// NewLabel make new *ui.Label
func NewLabel(w, h int, text string, face font.Face, color, bgColor color.Color, s int) *Label {
	r := image.Rect(0, 0, w, h)
	b := Box{r, bgColor, nil, nil, []Item{}, Callbacks{}, nil}
	l := &Label{b, text, face, color, s}
	l.Sub = l
	return l
}

// Reflesh updates internal *ebiten.Image
func (l *Label) Reflesh() {
	l.Box.Reflesh()
	_, h := l.Size()
	text.Draw(l.Image, l.Text, l.Face, 0, h-3, l.FontColor)
	// for _, r := range l.Text {
	// 	s := fmt.Sprint(l.Face.GlyphBounds(r))
	// 	fmt.Printf("%s:%s\n", string(r), s)
	// }
}

// String for fmt.Stringer interface
func (l *Label) String() string {
	p := fmt.Sprintf("%p", l)[7:11]
	return fmt.Sprintf("Label[%s]%s:%s", p, l.Text, l.Rect)
}
