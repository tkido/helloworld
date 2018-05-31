package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

type MouseManager struct {
	Downed *Downed
}

type Downed struct {
	Item  *Box
	Point image.Point
}

type MouseEventHandler interface {
	HandleMouseEvent(ev MouseEvent, origin image.Point, clip image.Rectangle) (handled bool, err error)
	SetCallback(tipe MouseEventType, c Callback)
}

// Callback is callback function on event
type Callback func(item Item)

type CallbacksToMounseEvents map[MouseEventType]Callback

func (cbs CallbacksToMounseEvents) SetCallback(t MouseEventType, c Callback) {
	cbs[t] = c
}

type MouseEvent struct {
	Type  MouseEventType
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

// IsCloseAsClick returns close enough the given two points (button downed and button upped) can be regarded as one click
func IsCloseAsClick(a, b image.Point) bool {
	sub := a.Sub(b)
	if sub.X*sub.X+sub.Y*sub.Y <= 16 {
		return true
	}
	return false
}
