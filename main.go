package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"strings"
	"time"

	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	imgSrc *ebiten.Image
	game   Game
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

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

	game = Game{false, false}
}

func update() error {
	return nil
}

func draw(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0x00, 0xff, 0xff, 0xff})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(0.1, 0.1)
	screen.DrawImage(imgSrc, opts)

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
		err = update()
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

func debugPrint(screen *ebiten.Image) (err error) {
	mx, my := ebiten.CursorPosition()
	buttons := []string{}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		buttons = append(buttons, "LEFT")
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		buttons = append(buttons, "RIGHT")
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		buttons = append(buttons, "MIDDLE")
	}

	pressed := []ebiten.Key{}
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			pressed = append(pressed, k)
		}
	}
	keyStrs := []string{}
	for _, p := range pressed {
		keyStrs = append(keyStrs, p.String())
	}

	sx, sy := ebiten.MonitorSize()

	const format = `FPS: %0.2f
mouse: (%d, %d) %v
keys: %s
IsCursorVisible: %v
DeviceScaleFactor: %v
IsFullscreen: %v
IsRunnableInBackground: %v
IsRunningSlowly: %v
IsWindowDecorated: %v
MonitorSize: (%d, %d)
ScreenScale: %0.2f
`
	msg := fmt.Sprintf(format,
		ebiten.CurrentFPS(),
		mx, my, buttons,
		strings.Join(keyStrs, ", "),
		ebiten.IsCursorVisible(),
		ebiten.DeviceScaleFactor(),
		ebiten.IsFullscreen(),
		ebiten.IsRunnableInBackground(),
		ebiten.IsRunningSlowly(),
		ebiten.IsWindowDecorated(),
		sx, sy,
		ebiten.ScreenScale(),
	)
	ebitenutil.DebugPrint(screen, msg)

	return
}
