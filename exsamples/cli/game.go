package main

// Game is Game
type Game struct {
	Screen Screen
	Name   string
}

// NewGame is NewGame
func NewGame() *Game {
	return &Game{
		Title,
		"",
	}
}

// Screen is Screen
type Screen int

// Screen enum
const (
	Title = iota
	Result
	Exit
)
