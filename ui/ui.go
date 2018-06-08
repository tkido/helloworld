package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Update ui
func Update(bg Item) {
	m.Now++
	// defered click event callback
	for i := 0; i < 3; i++ {
		click := LeftClick + EventType(i)
		if m.Clicked[i] != nil {
			if m.Now-m.Clicked[i].Frame > doubleClickInterval {
				m.Clicked[i].Item.Call(click)
				m.Clicked[i] = nil
			}
		}
	}
	// mouse control
	if ev, ok := m.getMouseEvent(); ok {
		bg.HandleMouseEvent(ev, image.ZP, bg.Rectangle())
		for k, v := range m.InItems {
			if v != m.Now {
				k.Call(MouseLeave)
				delete(m.InItems, k)
			}
		}
	}
}

// Draw ui
func Draw(screen *ebiten.Image, bg Item) {
	bg.Draw(screen, image.ZP, bg.Rectangle())
}
