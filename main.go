package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"time"

	"math/rand"

	"bitbucket.org/tkido/helloworld/vector"
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

var balls []*Ball

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
	if balls == nil {
		balls = []*Ball{}
		for i := 0; i < 20; i++ {
			r := rand.Float64()*20 + 5
			b := NewBall(
				r,
				vector.Vector{X: rand.Float64()*(screenWidth-2*r) + r, Y: rand.Float64()*(screenHeight-2*r) + r},
				vector.Vector{X: (rand.Float64() - 0.5) * 4, Y: (rand.Float64() - 0.5) * 4},
				imgSrc,
			)

			balls = append(balls, b)
		}
		return
	}
	for _, ball := range balls {
		ball.Update()
	}
	for _, b1 := range balls {
		for _, b2 := range balls {
			if b1 == b2 {
				continue
			}
			b1.IsCollision = b1.CheckCollision(b2)
			if b1.IsCollision {
				break
			}
		}
	}

	return
}

func draw(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0x00, 0xff, 0x00, 0xff})

	for _, ball := range balls {
		ball.Draw(screen)
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
		balls = nil
	}

	return
}

func main() {
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "Collision (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
