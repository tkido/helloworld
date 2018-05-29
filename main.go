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
	game    Game
	widgets []*ui.Box
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

func init() {
	game = Game{true, false}
	widgets = []*ui.Box{}

	box := ui.NewBox(100, 200, 50, 50, color.White)
	box2 := ui.NewBox(200, 300, 50, 50, color.Black)
	widgets = append(widgets, box, box2)
}

func control(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
	} else if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
	}

	if e, ok := ui.GetMouseEvent(); ok {
		if e.Point.In(screen.Bounds()) {
			for _, box := range widgets {
				if e.Point.In(box.Rect) {
					ok, err := box.HandleMouseEvent(e)
					if err != nil {
						return err
					}
					if ok {
						// end
					}
				}
			}
		}
	}
	return
}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0x00, 0xff, 0x00, 0xff})

	for _, box := range widgets {
		box.Draw(screen)
	}

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
