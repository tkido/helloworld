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
	screen := ui.NewBox(screenWidth, screenHeight, ui.Color("0f0"))
	menu := ui.NewNinepatch(400, 400, png, image.Rect(8, 8, 24, 24))
	screen.Add(10, 10, menu)

	result := ui.NewLabel(220, 30, "", 0, ui.Right, color.White, color.Black)
	screen.Add(420, 420, result)

	for i, s := range data {
		label := ui.NewLabel(400, 30, s, 0, ui.Left, ui.Color("0f0"), nil)
		label.SetCallback(ui.LeftClick, func(i int) func(item ui.Item) {
			return func(item ui.Item) {
				result.SetText(data[i])
			}
		}(i))
		label.SetCallback(ui.MouseOver, func(item ui.Item) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(20, 0)
			item.SetDIO(op)
		})
		label.SetCallback(ui.MouseOut, func(item ui.Item) {
			item.SetDIO(nil)
		})
		menu.Add(8, 8+i*30, label)
	}
	return screen
}
