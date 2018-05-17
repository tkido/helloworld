package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"time"

	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 512
	screenHeight = 512
)

var (
	imgSrc *ebiten.Image
	game   Game
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

// Drawable is drawable object
type Drawable interface {
	// SetScreen(*ebiten.Image) error
	// SetImage(*ebiten.Image) error
	// SetOptions(*ebiten.DrawImageOptions) error
	Draw() error
}

var ball *Ball

func init() {
	rand.Seed(time.Now().UnixNano())

	f, err := Assets.Open("/assets/nc35542.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	imgSrc, err = ebiten.NewImageFromImage(png, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	game = Game{true, false}
}

func update(screen *ebiten.Image) (err error) {
	if ball == nil {
		ball = NewBall(100, 100, 20, DrawableData{screen, imgSrc})
		return
	}
	ball.X++
	ball.Y++
	ball.R = ball.R + 0.2

	return
}

func draw(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0x00, 0xff, 0x00, 0xff})

	err = ball.Draw()
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

func mainLoop(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
	} else if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	} else if ebiten.IsKeyPressed(ebiten.KeyC) {
		game.IsRunning = false
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
		game.IsRunning = false
	}

	if game.IsRunning {
		err = update(screen)
		if err != nil {
			return
		}
	} else {

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

func main() {
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "Collision (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
