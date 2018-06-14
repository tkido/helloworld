package ui

import (
	"golang.org/x/image/font"
)

// FontManager manage font
type FontManager struct {
	Fonts   []font.Face
	Ascents []int
}

// AddFont add font face
func AddFont(face font.Face) {
	m.FontManager.Fonts = append(m.FontManager.Fonts, face)
	ascent := face.Metrics().Ascent.Floor()
	m.FontManager.Ascents = append(m.FontManager.Ascents, ascent)
}
