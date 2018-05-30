package main

import (
	"image/color"
	"log"

	"bitbucket.org/tkido/helloworld/ui"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	game Game
	bg   *ui.Box
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

func init() {
	game = Game{true, false}

	bg = ui.NewBox(0, 0, screenWidth, screenHeight, color.NRGBA{0x00, 0xff, 0x00, 0xff})
	bg.Add(ui.NewBox(100, 100, 200, 200, color.White))
	bg.Add(ui.NewBox(200, 200, 200, 200, color.Black))
}

func control(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
	} else if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
	}

	if e, ok := ui.GetMouseEvent(); ok {
		_, err := bg.HandleMouseEvent(e)
		if err != nil {
			log.Panicln(err)
		}
	}
	return
}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	bg.Draw(screen)

	return
}

func mainLoop(screen *ebiten.Image) (err error) {
	err = control(screen)
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
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "MouseEvent (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
