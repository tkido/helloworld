package main

import "bitbucket.org/tkido/helloworld/core/scene"

// Game is Game
type Game struct {
	Scene scene.Scene
	Name  string
}

// NewGame is NewGame
func NewGame() *Game {
	return &Game{
		scene.TITLE,
		"",
	}
}
