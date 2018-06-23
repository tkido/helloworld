package main

import (
	"io/ioutil"
	"log"

	"bitbucket.org/tkido/helloworld/assets"
	"bitbucket.org/tkido/helloworld/ui"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func init() {
	addFont("mplus-1p-regular.ttf", 24, 72, font.HintingFull)
	addFont("mplus-1p-bold.ttf", 12, 72, font.HintingFull)
	addFont("PixelMplus12-Regular.ttf", 24, 72, font.HintingFull)
}

func addFont(path string, size, dpi float64, hinting font.Hinting) {
	ttf, err := assets.FileSystem.Open(path)
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
		DPI:     dpi,
		Hinting: hinting,
	})
	ui.AddFont(face)
}
