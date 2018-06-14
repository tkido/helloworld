package ui

import (
	"github.com/hajimehoshi/ebiten"
)

// Manager is manager of internal status of ui
type Manager struct {
	Now int
	MouseManager
	KeyManager
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
			Pressed:        map[ebiten.Key]byte{},
			Callbacks:      map[ebiten.Key]KeyCallback{},
			RepeatInterval: 15,
		},
	}
}
