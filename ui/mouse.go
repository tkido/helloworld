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

type MouseEventType int

const (
	MouseEventMove MouseEventType = iota
	MouseEventDown
	MouseEventUp
	MouseEventDrag
	MouseEventDrop
)

var pressed [3]byte

func init() {
}

func GetMouseEvent() MouseEvent {
	for i := 0; i < 3; i++ {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton(i)) {
			pressed[i] = pressed[i]<<1 | 1
		} else {
			pressed[i] = pressed[i]<<1 | 0
		}
		if i == 0 {
			fmt.Printf("%08b\n", pressed[i])
		}
	}

	x, y := ebiten.CursorPosition()
	p := image.Point{x, y}

	e := MouseEvent{
		Type:  MouseEventMove,
		Point: p,
	}
	return e
}
