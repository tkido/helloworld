package main

import (
	"image"
	_ "image/png"
	"log"
	"time"

	"math/rand"

	"bitbucket.org/tkido/helloworld/perlin2d"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 320
)

var (
	game       Game
	noiseImage *image.RGBA
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

func init() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	game = Game{true, false}

	noiseImage = image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	const l = screenWidth * screenHeight
	var maxF, minF float64
	for i := 0; i < l; i++ {
		x := float64(i%screenWidth) / 10.0
		y := float64(i/screenWidth) / 10.0
		f := perlin2d.Fractal(x, y)
		if f < minF {
			minF = f
		} else if f > maxF {
			maxF = f
		}
	}
	// fmt.Printf("%f <= f <= %f", minF, maxF)
	rand.Seed(seed)
	for i := 0; i < l; i++ {
		x := float64(i%screenWidth) / 10.0
		y := float64(i/screenWidth) / 10.0
		f := perlin2d.Fractal(x, y)
		n := uint8((f - minF) / (maxF - minF) * 255)
		noiseImage.Pix[4*i] = n
		noiseImage.Pix[4*i+1] = n
		noiseImage.Pix[4*i+2] = n
		noiseImage.Pix[4*i+3] = 0xff
	}

}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	screen.ReplacePixels(noiseImage.Pix)

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
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "Perlin2D (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
