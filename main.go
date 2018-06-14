package main

import (
	_ "image/png"
	"log"

	"bitbucket.org/tkido/helloworld/ui"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 640
	screenHeight = 480

	dpi      = 72
	fontSize = 24
)

var (
	game       Game
	bg         *ui.Box
	normalFont font.Face
	count      int
)

// Game is status of game
type Game struct {
	IsDebugPrint bool
}

func init() {
	game = Game{false}
	bg = mouseScreen()
}

func control(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	}

	ui.Update(bg)
	return
}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	ui.Draw(screen, bg)
	return
}

func loop(screen *ebiten.Image) (err error) {
	err = control(screen)
	if err != nil {
		return
	}

	err = update(screen)
	if err != nil {
		return
	}

	if ebiten.IsRunningSlowly() {
		return
	}

	err = draw(screen)
	if err != nil {
		return
	}

	if game.IsDebugPrint {
		err = debugPrint(screen)
		if err != nil {
			return
		}
	}

	return
}

func main() {
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(loop, screenWidth, screenHeight, 1, "MouseEvent (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
