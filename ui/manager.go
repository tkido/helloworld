package ui

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

// Manager is manager of internal status of ui
type Manager struct {
	Now int
	MouseManager
	KeyManager
	FontManager
}

var m *Manager

func init() {
	m = &Manager{
		0,
		MouseManager{
			Downed:              [3]*MouseRecord{},
			Clicked:             [3]*MouseRecord{},
			InItems:             map[Item]int{},
			DoubleClickInterval: 15,
		},
		KeyManager{
			Status:         map[ebiten.Key]*KeyStatus{},
			RepeatInterval: 15,
		},
		FontManager{
			Fonts: []font.Face{},
		},
	}
}
