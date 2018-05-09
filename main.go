package main

import (
	"log"
	"os"

	"bitbucket.org/tkido/helloworld/core/godfather"
	tc "github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/mattn/go-runewidth"
)

var row = 0
var style = tc.StyleDefault

func main() {
	logf, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logf.Close()
	log.SetOutput(logf)

	s, err := tc.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}
	encoding.Register()
	if err = s.Init(); err != nil {
		log.Fatalln(err)
	}
	defer s.Fini()

	plain := tc.StyleDefault
	bold := style.Bold(true)

	s.SetStyle(tc.StyleDefault.Foreground(tc.ColorWhite).Background(tc.ColorBlack))
	s.Clear()

	quit := make(chan struct{})

	style = bold
	putln(s, "日本語表示のテスト"+s.CharacterSet())
	style = plain
	s.Show()

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tc.EventKey:
				switch ev.Key() {
				case tc.KeyEnter:
					s.Clear()
					row = 1
					for i := 0; i < 10; i++ {
						putln(s, godfather.Next())
					}
					s.Show()
				case tc.KeyEscape:
					close(quit)
					return
				case tc.KeyCtrlL:
					s.Sync()
				}
			case *tc.EventResize:
				s.Sync()
			}
		}
	}()

	<-quit

}

func putln(s tc.Screen, str string) {
	puts(s, style, 1, row, str)
	row++
}

func puts(s tc.Screen, style tc.Style, x, y int, str string) {
	i := 0
	var deferred []rune
	dwidth := 0
	for _, r := range str {
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
		case 1:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
}
