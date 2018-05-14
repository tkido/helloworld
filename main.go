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
	screenWidth  = 640 * 2
	screenHeight = 480 * 2
	blockSize    = 8
	dead         = 0
	live         = 1 << 4
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
			board[y][x] = rand.Intn(2) * live
		}
	}
}

func update(screen *ebiten.Image) error {
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
			case 17, 18, 21, 22, 23, 24, 25:
				board[y][x] = dead
			default:
				board[y][x] = c / live * live
			}
		}
	}

	if ebiten.IsRunningSlowly() {
		return nil
	}

	for y := range board {
		for x, c := range board[y] {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x)*blockSize, float64(y)*blockSize)
			opts.SourceRect = srcRects[c>>4]
			screen.DrawImage(imgSrc, opts)
		}
	}

	msg := fmt.Sprintf(`FPS: %0.2f`, ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Life (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
