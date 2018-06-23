package main

import (
	"image"
	"image/color"
	"log"

	"bitbucket.org/tkido/helloworld/assets"
	"bitbucket.org/tkido/helloworld/ui"
	"github.com/hajimehoshi/ebiten"
)

func onClick(i ui.Item) {
	log.Printf("%s %s", i, "clicked!!")
}
func onDoubleClick(i ui.Item) {
	log.Printf("%s %s", i, "double clicked!!!!")
}
func expand(i ui.Item) {
	w, h := i.Size()
	i.Resize(w+10, h+10)
}

func onMouseOn(i ui.Item) {
	log.Printf("%s %s", i, "MouseOn")
}
func onMouseIn(i ui.Item) {
	log.Printf("%s %s", i, "MouseIn")
}
func onMouseOver(i ui.Item) {
	log.Printf("%s %s", i, "MouseOver")
}
func onMouseOut(i ui.Item) {
	log.Printf("%s %s", i, "MouseOut")
}
func onMouseEnter(i ui.Item) {
	log.Printf("%s %s", i, "MouseEnter")
}
func onMouseLeave(i ui.Item) {
	log.Printf("%s %s", i, "MouseLeave")
}

func mouseScreen() *ui.Box {
	// images
	f, err := assets.FileSystem.Open("/assets/food_tenpura_ebiten.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	bg = ui.NewBox(screenWidth, screenHeight, color.NRGBA{0x00, 0xff, 0x00, 0xff})

	bg.SetCallback(ui.RightClick, onClick)
	bg.SetCallback(ui.RightDoubleClick, onDoubleClick)
	bg.Add(300, 200, ui.NewBox(200, 200, color.Black))

	box1 := ui.NewBox(200, 200, color.NRGBA{0x00, 0xff, 0xff, 0xff})
	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(2.0, 2.0)
	// op.ColorM.RotateHue(math.Pi)
	// w, h := box1.Size()
	// op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	// op.GeoM.Rotate(math.Pi / 5)
	// op.GeoM.Translate(float64(w)/2, float64(h)/2)
	box1.SetDIO(op)
	for i := -20; i <= 180; i += 100 {
		for j := -20; j <= 180; j += 100 {
			box := ui.NewBox(50, 50, color.NRGBA{0xff, 0x00, 0x00, 0xff})
			box.SetCallback(ui.LeftUp, expand)
			box1.Add(i, j, box)
		}
	}
	bg.Add(200, 100, box1)

	img := ui.NewImage(100, 100, png)
	img.SetCallback(ui.LeftClick, expand)

	box1.Add(-10, 120, img)

	label := ui.NewLabel(screenWidth, 30, "abcdefj", 0, ui.Left, color.White, color.Black)
	label2 := ui.NewLabel(screenWidth, 30, "テストですよ。", 0, ui.Right, color.White, color.Black)
	// label.SetCallback(ui.MouseOn, onMouseOn)
	// label.SetCallback(ui.MouseIn, onMouseIn)
	label.SetCallback(ui.MouseOut, onMouseOut)
	label.SetCallback(ui.MouseOver, onMouseOver)
	label.SetCallback(ui.MouseEnter, onMouseEnter)
	label.SetCallback(ui.MouseLeave, onMouseLeave)
	// bg.SetCallback(ui.MouseOn, onMouseOn)
	// bg.SetCallback(ui.MouseIn, onMouseIn)

	img.SetCallback(ui.RightClick, func(i ui.Item) {
		x, y := i.Position()
		i.Move(x+10, y)
	})

	bg.SetCallback(ui.MouseOut, onMouseOut)
	bg.SetCallback(ui.MouseOver, onMouseOver)
	bg.SetCallback(ui.MouseEnter, onMouseEnter)
	bg.SetCallback(ui.MouseLeave, onMouseLeave)
	bg.Add(10, 10, label)
	bg.Add(10, 40, label2)

	return bg
}
