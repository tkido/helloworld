package main

import (
	"fmt"
	"image/color"

	eb "github.com/hajimehoshi/ebiten"
	util "github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	if err := eb.Run(update, 640, 480, 1, "Hello world!"); err != nil {
		panic(err)
	}
}

var red = color.NRGBA{0xff, 0x00, 0x00, 0xff}

var i int
var square *eb.Image

func update(screen *eb.Image) error {
	screen.Fill(red)
	util.DebugPrint(screen, fmt.Sprintf("%d", i))
	i++

	if square == nil {
		square, _ = eb.NewImage(16, 16, eb.FilterNearest)
	}
	square.Fill(color.White)
	opts := &eb.DrawImageOptions{}
	opts.GeoM.Translate(32, 32)
	screen.DrawImage(square, opts)

	return nil
}
