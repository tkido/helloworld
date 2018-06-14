package ui

import (
	"log"

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
	log.Printf("%#v", face.Metrics().Ascent.Ceil())
	log.Printf("%#v", face.Metrics().Descent.Ceil())
	log.Printf("%#v", face.Metrics().Height.Ceil())
	m.FontManager.Ascents = append(m.FontManager.Ascents, ascent)
}
