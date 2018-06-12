package main

import (
	"image"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"

	"bitbucket.org/tkido/helloworld/ui"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 640
	screenHeight = 480

	dpi      = 72
	fontSize = 24
)

var (
	game       Game
	bg         *ui.Box
	normalFont font.Face
	count      int
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

// Game is status of game
type Game struct {
	IsRunning, IsDebugPrint bool
}

func init() {
	// images
	f, err := Assets.Open("/assets/food_tenpura_ebiten.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	// fonts
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

	normalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	game = Game{true, false}

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
	box1.DrawImageOptions = op
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

	label := ui.NewLabel(screenWidth, 30, ".fjあいうアイウ愛飢男■★◆Ａｊｆ", normalFont, color.White, color.Black, 24)
	label2 := ui.NewLabel(screenWidth, 30, ".fjあいうアイウ愛飢男■★◆Ａｊｆ", normalFont, color.White, color.Black, 24)
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
}

func control(screen *ebiten.Image) (err error) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.IsRunning = !game.IsRunning
	} else if ebiten.IsKeyPressed(ebiten.KeyF4) {
		game.IsDebugPrint = !game.IsDebugPrint
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {

	}

	ui.Update(bg)
	return
}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	ui.Draw(screen, bg)
	return
}

func mainLoop(screen *ebiten.Image) (err error) {
	err = control(screen)
	if err != nil {
		return
	}

	if game.IsRunning {
		err = update(screen)
		if err != nil {
			return
		}
	}

	if ebiten.IsRunningSlowly() {
		return
	}

	err = draw(screen)
	if err != nil {
		return
	}

	if game.IsDebugPrint {
		err = debugPrint(screen)
		if err != nil {
			return
		}
	}

	return
}

func main() {
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(mainLoop, screenWidth, screenHeight, 1, "MouseEvent (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
