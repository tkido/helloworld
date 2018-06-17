package main

import (
	"io/ioutil"
	"log"

	"bitbucket.org/tkido/helloworld/ui"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	dpi      = 72
	fontSize = 24
)

func init() {
	addFont("/assets/PixelMplus12-Regular.ttf", 24, 72, font.HintingFull)
	addFont("/assets/mplus-1p-regular.ttf", 24, 72, font.HintingFull)
}

func addFont(path string, size, DPI float64, hinting font.Hinting) {
	ttf, err := Assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer ttf.Close()
	bs, err := ioutil.ReadAll(ttf)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(bs)
	if err != nil {
		log.Fatal(err)
	}
	face := truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     DPI,
		Hinting: hinting,
	})
	ui.AddFont(face)
}