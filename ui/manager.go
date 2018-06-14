package ui

import (
	"image"

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

// MouseManager manage status of mouse for ui
type MouseManager struct {
	pressed             [3]byte
	last                MouseEvent
	Downed, Clicked     [3]*MouseRecord
	OnItem              Item
	InItems             map[Item]int
	DoubleClickInterval int // max interval recognized as DoubleClick. Unit is frame(1/60 second)
}

// GetMouseEvent make new mouse event
func (m *MouseManager) getMouseEvent() (e MouseEvent, updated bool) {
	moves := [3]MouseButtonMove{}
	for i := 0; i < 3; i++ {
		var pressed byte
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton(i)) {
			pressed = 1
		}
		m.pressed[i] = m.pressed[i]<<1 | pressed
		moves[i] = MouseButtonMove(m.pressed[i] & 3)
	}
	x, y := ebiten.CursorPosition()
	p := image.Point{x, y}
	e = MouseEvent{moves, p}
	if e != m.last {
		m.last = e
		return e, true
	}
	return e, false
}
