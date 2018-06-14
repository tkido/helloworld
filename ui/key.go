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
			k.Callbacks[key]()
		}
	}
}

// SetCallback set callback function for key. set nil means delete.
func SetCallback(key ebiten.Key, cb KeyCallback) {
	k := m.KeyManager
	if cb == nil {
		delete(k.Pressed, key)
		delete(k.Callbacks, key)
		return
	}
	if _, ok := k.Pressed[key]; !ok {
		k.Pressed[key] = 0
	}
	k.Callbacks[key] = cb
}
