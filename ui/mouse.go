package ui

import (
	"fmt"
	"image"
)

// doubleClickInterval as frame(1/60 second)
const doubleClickInterval = 15

// MouseButtonMove is move of mouse button
type MouseButtonMove int

// MouseButtonMove definition
const (
	None    MouseButtonMove = 0
	Down                    = 1
	Up                      = 2
	PRESSED                 = 3
)

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
