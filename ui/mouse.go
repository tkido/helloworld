package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// doubleClickInterval as frame(1/60 second)
const doubleClickInterval = 15

// MouseButtonMove is move of mouse button
type MouseButtonMove int

// MouseButtonMove definition
const (
	Down MouseButtonMove = 1
	Up                   = 2
)

// MouseManager manage status of mouse for ui
type MouseManager struct {
	pressed         [3]byte
	last            MouseEvent
	Downed, Clicked [3]*MouseRecord
	Overed          Item
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

// MouseRecord is record of mouse move and event
type MouseRecord struct {
	Item  *Box
	Point image.Point
	Frame int
}

// MouseEventHandler is
type MouseEventHandler interface {
	HandleMouseEvent(ev MouseEvent, origin image.Point, clip image.Rectangle) (handled bool, err error)
	SetCallback(tipe EventType, c Callback)
}

// MouseEvent is event about mouse action
type MouseEvent struct {
	Moves [3]MouseButtonMove
	Point image.Point
}

// String for fmt.Stringer interface
func (e MouseEvent) String() string {
	return fmt.Sprintf("%v%s", e.Moves, e.Point)
}

// EventType is type of all UI event
type EventType int

// MouseEvents
const (
	LeftDown EventType = iota
	RightDown
	MiddleDown
	LeftUp
	RightUp
	MiddleUp
	LeftClick
	RightClick
	MiddleClick
	LeftDoubleClick
	RightDoubleClick
	MiddleDoubleClick
	MouseOver
	MouseOut
	MouseEnter
	MouseLeave
)
