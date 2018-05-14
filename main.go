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
	board    []int
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

	board = make([]int, w*h, w*h)
	deltas = []int{-1, 0, 1}

	for i := range board {
		board[i] = rand.Intn(2) * live
	}
}

func update(screen *ebiten.Image) error {
	for i, c := range board {
		if c < live {
			continue
		}
		x, y := i%w, i/w
		for _, dy := range deltas {
			for _, dx := range deltas {
				board[((y+dy+h)%h)*w+(x+dx+w)%w]++
			}
		}
	}
	for i, c := range board {
		switch c {
		case 3:
			board[i] = live
		case 17, 18, 21, 22, 23, 24, 25:
			board[i] = dead
		default:
			board[i] = c / live * live
		}
	}

	if ebiten.IsRunningSlowly() {
		return nil
	}

	for i := range board {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(i%w)*blockSize, float64(i/w)*blockSize)
		opts.SourceRect = srcRects[board[i]>>4]
		screen.DrawImage(imgSrc, opts)
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
