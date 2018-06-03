package ui

import "github.com/hajimehoshi/ebiten"

// Update ui
func Update(bg Item) {
	m.Now++
	// defered click event callback
	if m.Clicked != nil {
		if m.Now-m.Clicked.Frame > doubleClickInterval {
			if c, ok := m.Clicked.Item.Callbacks[MouseClick]; ok {
				c(m.Clicked.Item.Sub)
			}
			m.Clicked = nil
		}
	}

	if ev, ok := GetMouseEvent(); ok {
		rect := bg.Rectangle()
		bg.HandleMouseEvent(ev, rect.Min, rect)
	}
}

// Draw ui
func Draw(screen *ebiten.Image, bg Item) {
	rect := bg.Rectangle()
	bg.Draw(screen, rect.Min, rect)
}
