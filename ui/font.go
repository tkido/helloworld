package ui

import (
	"golang.org/x/image/font"
)

// Font is data of font
type Font struct {
	Face   font.Face
	Ascent int
}

// FontManager manage font
type FontManager struct {
	Fonts []Font
}

// AddFont add font face
func AddFont(face font.Face) {
	ascent := face.Metrics().Ascent.Ceil()
	data := Font{face, ascent}
	m.FontManager.Fonts = append(m.FontManager.Fonts, data)
}
