package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"strings"
	"time"

	"bitbucket.org/tkido/helloworld/core/godfather"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	sampleText      = `The quick brown fox jumps over the lazy dog.`
	mplusNormalFont font.Face
	mplusBigFont    font.Face
	counter         = 0
	kanjiText       = []rune{}
	kanjiTextColor  color.RGBA
)

var jaKanjis = []rune{}

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	// Change the text color for each second.
	if counter%ebiten.FPS == 0 {
		kanjiText = []rune{}
		for j := 0; j < 10; j++ {
			// kanjiText = append(kanjiText, jaKanjis[rand.Intn(len(jaKanjis))])
			for _, r := range godfather.Next() {
				kanjiText = append(kanjiText, r)
			}
			kanjiText = append(kanjiText, '\n')
		}

		kanjiTextColor.R = 0x80 + uint8(rand.Intn(0x7f))
		kanjiTextColor.G = 0x80 + uint8(rand.Intn(0x7f))
		kanjiTextColor.B = 0x80 + uint8(rand.Intn(0x7f))
		kanjiTextColor.A = 0xff
	}
	counter++

	if ebiten.IsRunningSlowly() {
		return nil
	}

	f, err := Assets.Open("/assets/block00.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	block, err := ebiten.NewImageFromImage(png, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(100, 100)
	screen.DrawImage(block, opts)
	opts.GeoM.Translate(200, 200)
	r := image.Rect(8, 0, 16, 8)
	opts.SourceRect = &r
	screen.DrawImage(block, opts)

	const x = 20

	// Draw info
	msg := fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS())
	text.Draw(screen, msg, mplusNormalFont, x, 40, color.White)

	// Draw the sample text
	text.Draw(screen, sampleText, mplusNormalFont, x, 80, color.White)

	// Draw Kanji text lines
	for i, line := range strings.Split(string(kanjiText), "\n") {
		text.Draw(screen, line, mplusNormalFont, x, 160+30*i, kanjiTextColor)
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Font (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
