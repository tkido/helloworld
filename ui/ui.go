package ui

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
)

// Update ui
func Update(bg Item) {
	m.Now++
	m.KeyManager.KeyEvent()
	// mouse control
	if ev, ok := m.getMouseEvent(); ok {
		if handled := bg.HandleMouseEvent(ev, image.ZP, image.Rect(0, 0, math.MaxInt64, math.MaxInt32)); !handled {
			if m.OnItem != nil {
				m.OnItem.Call(MouseOut)
				m.OnItem = nil
			}
		}
		for k, v := range m.InItems {
			if v != m.Now {
				k.Call(MouseLeave)
				delete(m.InItems, k)
			}
		}
	}
	// defered click event callback
	for i := 0; i < 3; i++ {
		if m.Clicked[i] != nil && m.Now-m.Clicked[i].Frame > m.DoubleClickInterval {
			click := LeftClick + EventType(i)
			m.Clicked[i].Item.Call(click)
			m.Clicked[i] = nil
		}
	}
}

// Draw ui
func Draw(target *ebiten.Image, bg Item) {
	bg.Draw(target)
}
