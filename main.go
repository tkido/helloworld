package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	runewidth "github.com/mattn/go-runewidth"
)

var defStyle tcell.Style

func main() {
	encoding.Register()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer s.Fini()
	defStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.Clear()

	w, h := s.Size()
	white := tcell.StyleDefault.
		Foreground(tcell.ColorWhite)
	w, h = s.Size()
	emitStr(s, 0, 0, white, fmt.Sprint(w))
	emitStr(s, 0, 1, white, fmt.Sprint(h))
	puts(s, white, 0, 2, "テストですよ")
	s.Show()

MAINLOOP:
	for {

		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			s.SetContent(w-1, h-1, 'R', nil, white)
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				break MAINLOOP
			}
		}
	}
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func puts(s tcell.Screen, style tcell.Style, x, y int, str string) {
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
