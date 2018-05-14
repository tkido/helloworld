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
	board    [][]int
	deltas   []int
)

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

	board = make([][]int, h, h)
	deltas = []int{-1, 0, 1}

	for y := range board {
		board[y] = make([]int, w, w)
		for x := range board[y] {
			if rand.Float64() < 0.3 {
				board[y][x] = live
			}
		}
	}
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

	msg := fmt.Sprintf(`FPS: %0.2f`, ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	return
}

func updateScreen(screen *ebiten.Image) (err error) {
	err = update()
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
	return
}

func main() {
	if err := ebiten.Run(updateScreen, screenWidth, screenHeight, 1, "Life (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
