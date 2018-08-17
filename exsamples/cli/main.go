package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("START")
	loop()
	log.Println("END")
}

func loop() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		str := s.Text()
		log.Print(str)
		switch str {
		case "exit":
			log.Print("hogehoge")
			break
		}
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}

// Game is Game
type Game struct {
	Time time.Time
}

// NewGame is NewGame
func NewGame() *Game {
	return &Game{time.Now()}
}
