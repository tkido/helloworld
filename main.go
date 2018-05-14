package main

import (
	"image"
	_ "image/png"
	"log"
	"time"

	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
	blockSize    = 8
	dead         = 0
	live         = 1 << 4
)

var (
	imgSrc   *ebiten.Image
	srcRects []*image.Rectangle
	board    []int
	moores   []int
)

const (
	w = screenWidth / blockSize
	h = screenHeight / blockSize
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// イメージソースの準備
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

	for i := range board {
		board[i] = rand.Intn(2) * live
	}
}

func update(screen *ebiten.Image) error {
	// screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	if ebiten.IsRunningSlowly() {
		return nil
	}
	for i := range board {
		// if c == dead {
		// 	continue
		// }
		board[i] = rand.Intn(2) * live
	}
	for i, c := range board {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(i%w)*blockSize, float64(i/w)*blockSize)
		opts.SourceRect = srcRects[c>>4]
		screen.DrawImage(imgSrc, opts)
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Life (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
