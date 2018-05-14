package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
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

// Board in matrix of cell
type Board [][]int

// NewBoard get Board
func NewBoard() Board {
	b := make([][]int, h, h)
	for y := range b {
		b[y] = make([]int, w, w)
	}
	return b
}

// Clear all cells
func (b Board) Clear() {
	for y := range b {
		for x := range b[y] {
			b[y][x] = dead
		}
	}
}

// Rand randomize all cells
func (b Board) Rand() {
	for y := range b {
		for x := range b[y] {
			if rand.Float64() < 0.3 {
				b[y][x] = live
			} else {
				b[y][x] = dead
			}
		}
	}
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
		msg := fmt.Sprintf(`FPS: %0.2f`, ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, msg)
	}

	return
}

func mainLoop(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
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
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "Life (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
