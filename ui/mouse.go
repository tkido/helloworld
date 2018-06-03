package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// MouseManager manage status of mouse for ui
type MouseManager struct {
	Downed, Clicked *MouseRecord
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
	Type  EventType
	Point image.Point
}

// String for fmt.Stringer interface
func (ev MouseEvent) String() string {
	var name string
	switch ev.Type {
	case MouseMove:
		name = "Move"
	case MouseDown:
		name = "Down"
	case MouseUp:
		name = "Up"
	case MouseDrag:
		name = "Drag"
	case MouseDrop:
		name = "Drop"
	}
	return fmt.Sprintf("%s%s", name, ev.Point)
}

// EventType is type of all UI event
type EventType int

const (
	MouseMove EventType = iota
	MouseDown
	MouseUp
	MouseDrag
	MouseDrop
	MouseOver
	MouseLeave
	MouseClick
	MouseDoubleClick
)

var pressed [3]byte
var last MouseEvent

// doubleClickInterval as frame(1/60 second)
const doubleClickInterval = 15

// GetMouseEvent make new mouse event
func GetMouseEvent() (e MouseEvent, updated bool) {

	for i := 0; i < 3; i++ {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton(i)) {
			pressed[i] = pressed[i]<<1 | 1
		} else {
			pressed[i] = pressed[i]<<1 | 0
		}
	}
	tipe := EventType(pressed[0] & 3)

	x, y := ebiten.CursorPosition()
	p := image.Point{x, y}
	e = MouseEvent{tipe, p}

	if e != last {
		last = e
		return e, true
	}
	return e, false
}
