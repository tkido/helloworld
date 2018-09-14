package main

import (
	"image/color"
	_ "image/png"
	"log"

	"bitbucket.org/tkido/helloworld/ui"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	game  Game
	bg    *ui.Box
	count int
)

// Game is status of game
type Game struct {
	IsDebugPrint bool
}

func init() {
	game = Game{false}
	bg = ui.NewBox(screenWidth, screenHeight, color.NRGBA{0x00, 0xff, 0x00, 0xff})
	menu := menuScreen()
	mainMenu := mainMenu()
	bg.Add(0, 0, mainMenu)
	bg.Add(0, 0, menu)
	ui.SetCallback(ebiten.KeyF4, func() {
		game.IsDebugPrint = !game.IsDebugPrint
	})
	ui.SetCallback(ebiten.KeyF5, func() {
		ebiten.SetScreenScale(2)
	})
}

func control(screen *ebiten.Image) (err error) {
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
