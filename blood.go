package main

import (
	_ "image/jpeg"

	"bitbucket.org/tkido/helloworld/ui"
)

func bloodScreen() *ui.Box {
	face1 := ui.NewImage(96*2, 120*2, loadImage("face/face0038.jpg"))
	face2 := ui.NewImage(96*2, 120*2, loadImage("face/face0108.jpg"))

	msgWin := ui.NewImage(640, 151, loadImage("window/box_red_name.png"))
	label := ui.NewLabel(600, 140, "ブラド", 0, ui.Left, ui.Color("fff"), nil)
	msg := ui.NewLabel(600, 140, "さて、どうやって料理してやろうか……。", 0, ui.Left, ui.Color("fff"), nil)

	screen := ui.NewBox(screenWidth, screenHeight, ui.Color("900"))
	screen.Add(100, 50, face1)
	screen.Add(300, 50, face2)
	screen.Add(0, 480-151, msgWin)
	screen.Add(20, 480-150, label)
	screen.Add(10, 480-120, msg)
	return screen
}
