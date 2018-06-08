package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Manager is manager of internal status of ui
type Manager struct {
	Now int
	MouseManager
}

var m *Manager

func init() {
	m = &Manager{
		0,
		MouseManager{
			Downed:  [3]*MouseRecord{},
			Clicked: [3]*MouseRecord{},
			InItems: map[Item]int{},
		},
	}
}

// MouseManager manage status of mouse for ui
type MouseManager struct {
	pressed         [3]byte
	last            MouseEvent
	Downed, Clicked [3]*MouseRecord
	OnItem          Item
	InItems         map[Item]int
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
		move := MouseButtonMove(m.pressed[i] & 3)
		switch move {
		case Down, Up:
		default:
			move = None
		}
		moves[i] = move
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
