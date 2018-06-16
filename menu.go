package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"

	"bitbucket.org/tkido/helloworld/ui"
)

func menuScreen() *ui.Box {
	// images
	f, err := Assets.Open("/assets/ninepatch.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	data := []string{"バナナ", "りんご", "オレンジ", "とうもろこし", "焼き肉", "北京ダック", "豚の丸焼き"}
	bg = ui.NewBox(screenWidth, screenHeight, color.NRGBA{0x00, 0xff, 0x00, 0xff})
	// menu := ui.NewBox(400, 400, color.NRGBA{0x00, 0xff, 0xff, 0xff})
	menu := ui.NewNinepatch(400, 400, png)
	bg.Add(10, 10, menu)

	result := ui.NewLabel(240, 30, "", 0, color.White, color.Black)
	bg.Add(420, 420, result)

	img := ui.NewImage(32, 32, png)
	nine := ui.NewNinepatch(128, 128, png)
	bg.Add(440, 100, img)
	bg.Add(440, 150, nine)

	for i, s := range data {
		label := ui.NewLabel(400, 30, s, 0, color.NRGBA{0x00, 0xff, 0x00, 0xff}, nil)
		onSelect := func(i int) func(item ui.Item) {
			return func(item ui.Item) {
				result.SetText(data[i])
			}
		}(i)
		label.SetCallback(ui.LeftClick, onSelect)
		label.SetCallback(ui.MouseOver, func(item ui.Item) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(20, 0)
			item.SetDIO(op)
			// x, y := item.Position()
			// item.Move(x+20, y)
		})
		label.SetCallback(ui.MouseOut, func(item ui.Item) {
			item.SetDIO(nil)
			// x, y := item.Position()
			// item.Move(x-20, y)
		})
		menu.Add(8, 8+i*30, label)
	}
	return bg
}
