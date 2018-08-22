package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"bitbucket.org/tkido/helloworld/core/scene"
)

func main() {
	g := NewGame()
LOOP:
	for {
		switch g.Scene {
		case scene.TITLE:
			g.title()
		case scene.NEWGAME:
			g.newGame()
		case scene.EXIT:
			fmt.Println("ゲームを終了します。")
			break LOOP
		}
	}
}

func (g *Game) title() {
	fmt.Println("Vampire Hunter Version 0.0.0.0")
	c := NewChoice()
	c.Add(scene.NEWGAME, "New Game")
	c.Add(scene.LOADGAME, "Load Game")
	c.Add(scene.EXIT, "Exit Game")
	r := c.Do()
	g.Scene = scene.Scene(r)
}

func (g *Game) newGame() {
	const (
		NO = iota
		YES
	)
	for g.Name == "" {
		name := getName()
		msg := fmt.Sprintf("あなたの名前は%sです。", name)
		fmt.Println(msg)
		fmt.Println("よろしいですか？")

		c := NewChoice()
		c.Add(YES, "はい")
		c.Add(NO, "いいえ")
		switch c.Do() {
		case YES:
			g.Name = name
		}
	}
	g.Scene = scene.EXIT
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
