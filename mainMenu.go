package main

import (
	"image/color"

	"bitbucket.org/tkido/helloworld/ui"
	"github.com/hajimehoshi/ebiten"
)

func mainMenu() *ui.Box {
	data := []string{"マウス", "キー", "メニュー", "ダイアログ"}
	menu := ui.NewBox(640, 30, color.NRGBA{0x99, 0x99, 0x99, 0xff})
	ui.SetCallback(ebiten.KeyEscape, func() {
		if menu.IsVisible() {
			menu.Hide()
			return
		}
		menu.Show()
	})
	for i, s := range data {
		label := ui.NewLabel(100, 30, s, 0, ui.Center, color.Black, nil)
		label.SetCallback(ui.LeftClick, func(i int) func(item ui.Item) {
			return func(item ui.Item) {
				// result.SetText(data[i])
			}
		}(i))
		label.SetCallback(ui.MouseOver, func(item ui.Item) {
			// op := &ebiten.DrawImageOptions{}
			// op.GeoM.Translate(20, 0)
			// item.SetDIO(op)
		})
		label.SetCallback(ui.MouseOut, func(item ui.Item) {
			// item.SetDIO(nil)
		})
		menu.Add(i*120, 0, label)
	}
	return menu
}
