package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Item is ebiten UI item
type Item interface {
	Draw(screen *ebiten.Image, origin image.Point, clip image.Rectangle)
	Reflesh()
	Move(x, y int)
	Resize(w, h int)
	Size() (w, h int)
	Add(x, y int, item Item)
	HandleMouseEvent(ev MouseEvent, origin image.Point, clip image.Rectangle) (handled bool)
}
