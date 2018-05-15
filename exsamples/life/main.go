package main

import (
	"fmt"
	"image"
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
	blockSize    = 8
	dead         = 0
	live         = 9
)

var (
	imgSrc   *ebiten.Image
	srcRects []*image.Rectangle
	board    Board
	deltas   []int
	game     Game
)

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

const (
	w = screenWidth / blockSize
	h = screenHeight / blockSize
)

func init() {
	rand.Seed(time.Now().UnixNano())

	f, err := Assets.Open("/assets/block00.png")
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

	deadRect := image.Rect(0, 0, blockSize, blockSize)
	liveRect := image.Rect(blockSize, 0, blockSize*2, blockSize)
	srcRects = []*image.Rectangle{&deadRect, &liveRect}
	deltas = []int{-1, 0, 1}

	game = Game{false, false}
	board = NewBoard()
	board.Rand()

}

func update() error {
	for y := range board {
		for x, c := range board[y] {
			if c < live {
				continue
			}
			for _, dy := range deltas {
				for _, dx := range deltas {
					board[(y+dy+h)%h][(x+dx+w)%w]++
				}
			}
		}
	}
	for y := range board {
		for x, c := range board[y] {
			switch c {
			case 3:
				board[y][x] = live
			case 10, 11, 14, 15, 16, 17, 18:
				board[y][x] = dead
			default:
				board[y][x] = c / live * live
			}
		}
	}
	return nil
}

func draw(screen *ebiten.Image) (err error) {
	for y := range board {
		for x, c := range board[y] {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x)*blockSize, float64(y)*blockSize)
			opts.SourceRect = srcRects[c/live]
			screen.DrawImage(imgSrc, opts)
		}
	}

	if game.IsDebugPrint {
		err = debugPrint(screen)
		if err != nil {
			return
		}
	}

	return
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

func mainLoop(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
	} else if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	} else if ebiten.IsKeyPressed(ebiten.KeyC) {
		board.Clear()
		game.IsRunning = false
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
		board.Rand()
		game.IsRunning = false
	}

	if game.IsRunning {
		err = update()
		if err != nil {
			return
		}
	} else {
		mx, my := ebiten.CursorPosition()
		x, y := mx/blockSize, my/blockSize
		if 0 <= x && x < w && 0 <= y && y < h {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				board[y][x] = live
			} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
				board[y][x] = dead
			}
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

func main() {
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "Life (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
