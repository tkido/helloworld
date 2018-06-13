package ui

// Callback is callback function on event
type Callback interface{}

// Callbacks keep callback functions
type Callbacks map[EventType]Callback

// SetCallback set callback function to item. When set nil, it means delete
func (cs Callbacks) SetCallback(t EventType, c Callback) {
	if c == nil {
		delete(cs, t)
	} else {
		cs[t] = c
	}
}
