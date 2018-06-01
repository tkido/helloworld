package ui

// Manager is manager of internal status of ui
type Manager struct {
	Now int
	MouseManager
}

var m *Manager

func init() {
	m = &Manager{
		0,
		MouseManager{
			Downed:  nil,
			Clicked: nil,
		},
	}
}
