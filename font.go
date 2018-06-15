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
	ttf, err := Assets.Open("/assets/PixelMplus12-Regular.ttf")
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
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	ui.AddFont(face)
}
