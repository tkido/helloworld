package ui

import (
	"github.com/hajimehoshi/ebiten"
)

// KeyCallback is callback for key
type KeyCallback func()

// KeyStatus is status of key
type KeyStatus struct {
	Pressed  byte
	Callback KeyCallback
}

// KeyManager manage status of key
type KeyManager struct {
	Status         map[ebiten.Key]*KeyStatus
	RepeatInterval int
}

// KeyEvent call event
func (k *KeyManager) KeyEvent() {
	for key, s := range k.Status {
		var pressed byte
		if ebiten.IsKeyPressed(key) {
			pressed = 1
		}
		s.Pressed = s.Pressed<<1 | pressed
		if s.Pressed&3 == 1 {
			s.Callback()
		}
	}
}

// SetCallback set callback function for key. set nil means delete.
func SetCallback(key ebiten.Key, cb KeyCallback) {
	k := m.KeyManager
	if cb == nil {
		delete(k.Status, key)
		return
	}
	if s, ok := k.Status[key]; ok {
		s.Callback = cb
	} else {
		k.Status[key] = &KeyStatus{0, cb}
	}
}
