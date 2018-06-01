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
	game Game
	bg   *ui.Box
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

func init() {
	game = Game{true, false}

	bg = ui.NewBox(screenWidth, screenHeight, color.NRGBA{0x00, 0xff, 0x00, 0xff})
	bg.Add(200, 200, ui.NewBox(200, 200, color.Black))

	box1 := ui.NewBox(200, 200, color.White)
	for i := -20; i <= 180; i += 100 {
		for j := -20; j <= 180; j += 100 {
			box1.Add(i, j, ui.NewBox(50, 50, color.NRGBA{0xff, 0x00, 0x00, 0xff}))
		}
	}
	bg.Add(100, 100, box1)
}

func control(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
	} else if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
	}

	if ev, ok := ui.GetMouseEvent(); ok {
		bg.HandleMouseEvent(ev, screen.Bounds().Min, screen.Bounds())
	}
	return
}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	bg.Draw(screen, screen.Bounds().Min, screen.Bounds())
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
