package ui

import (
	"github.com/hajimehoshi/ebiten"
)

// KeyCallback is callback for key
type KeyCallback func()

// KeyManager manage status of key
type KeyManager struct {
	Pressed        map[ebiten.Key]byte
	Callbacks      map[ebiten.Key]KeyCallback
	RepeatInterval int
}

// KeyEvent call event
func (k *KeyManager) KeyEvent() {
	for key, b := range k.Pressed {
		var pressed byte
		if ebiten.IsKeyPressed(key) {
			pressed = 1
		}
		b = b<<1 | pressed
		k.Pressed[key] = b
		if b&3 == 1 {
			if cb, ok := k.Callbacks[key]; ok {
				cb()
			}
		}
	}
}
