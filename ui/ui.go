package ui

var m *UIManager

func init() {
	m = &UIManager{
		MouseManager{
			Downed: nil,
		},
	}
}
