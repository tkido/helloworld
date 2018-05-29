package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

type MouseEventHandler interface {
	HandleMouseEvent(e MouseEvent) (handled bool, err error)
}

type MouseEvent struct {
	Type  MouseEventType
	Point image.Point
}

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

type MouseEventType int

const (
	MouseMove MouseEventType = iota
	MouseDown
	MouseUp
	MouseDrag
	MouseDrop
	MouseOver
	MouseLeave
)

var pressed [3]byte
var last MouseEvent

func init() {

}

func GetMouseEvent() (e MouseEvent, updated bool) {
	for i := 0; i < 3; i++ {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton(i)) {
			pressed[i] = pressed[i]<<1 | 1
		} else {
			pressed[i] = pressed[i]<<1 | 0
		}
	}
	tipe := MouseEventType(pressed[0] & 3)

	x, y := ebiten.CursorPosition()
	p := image.Point{x, y}
	e = MouseEvent{tipe, p}

	if e != last {
		last = e
		return e, true
	}
	return e, false
}
