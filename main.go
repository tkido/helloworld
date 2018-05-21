package main

import (
	"image/color"
	_ "image/png"
	"log"
	"time"

	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640 / 2
	screenHeight = 480 / 2
)

var (
	game Game
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
	game = Game{true, false}
}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0x00, 0xff, 0x00, 0xff})

	if game.IsDebugPrint {
		err = debugPrint(screen)
		if err != nil {
			return
		}
	}

	return
}

func mainLoop(screen *ebiten.Image) (err error) {
	err = control()
	if err != nil {
		return
	}

	if game.IsRunning {
		err = update(screen)
		if err != nil {
			return
		}
	}

	if ebiten.IsRunningSlowly() {
		return
	}

	err = draw(screen)
	if err != nil {
		return
	}

	return
}

func control() (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
	} else if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
	}

	return
}

func main() {
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "Collision (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
