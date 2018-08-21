package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	g := NewGame()
LOOP:
	for {
		switch g.Screen {
		case Title:
			g.title()
		case Result:
			g.result()
		case Exit:
			break LOOP
		}
	}
}

func (g *Game) title() {
	fmt.Println("Vampire Hunter")
	g.Name = getName()
	g.Screen = Result
}

func (g *Game) result() {
	msg := fmt.Sprintf("あなたの名前は%sです。", g.Name)
	fmt.Println(msg)
	g.Screen = Exit
}

func getName() string {
	name := ""
	for name == "" {
		fmt.Println("名前を入力してください。")
		s := bufio.NewScanner(os.Stdin)
		if s.Scan() {
			name = s.Text()
		}
		if s.Err() != nil {
			log.Fatal(s.Err())
		}
	}
	return name
}
