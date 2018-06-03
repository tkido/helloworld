package ui

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
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
func NewLabel(w, h int, text string, face font.Face, c color.Color, s int) *Label {
	r := image.Rect(0, 0, w, h)
	b := Box{r, nil, nil, nil, []Item{}, Callbacks{}, nil}
	l := &Label{b, text, face, c, s}
	l.Box.Super = l
	return l
}

// Reflesh updates internal *ebiten.Image
func (l *Label) Reflesh() {
	fmt.Println("label Reflesh")
	w, h := l.Size()
	l.Image, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	text.Draw(l.Image, "テスト", l.Face, 0, 40, l.FontColor)
}

// String for fmt.Stringer interface
func (l *Label) String() string {
	p := fmt.Sprintf("%p", l)[7:11]
	return fmt.Sprintf("Label[%s]%s:%s", p, l.Text, l.Box.Rect)
}
