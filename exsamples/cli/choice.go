package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Choice is choice
type Choice struct {
	List []Alt
	Map  map[int]struct{}
}

// Alt is Alternative for Choice
type Alt struct {
	ID    int
	Label string
}

// NewChoice returns new Choice
func NewChoice() *Choice {
	return &Choice{
		[]Alt{},
		map[int]struct{}{},
	}
}

// Add add alt to Choice
func (c *Choice) Add(id int, label string) {
	a := Alt{id, label}
	if _, ok := c.Map[a.ID]; !ok {
		c.List = append(c.List, a)
		c.Map[a.ID] = struct{}{}
	}
}

// Do is Do
func (c *Choice) Do() int {
	if len(c.List) == 0 {
		return -1
	}
	format := "%2d: %s\n"
	buf := bytes.Buffer{}
	for i, a := range c.List {
		s := fmt.Sprintf(format, i, a.Label)
		buf.WriteString(s)
	}
	msg := buf.String()
	for {
		fmt.Println(msg)
		s := bufio.NewScanner(os.Stdin)
		if s.Scan() {
			tmp := s.Text()
			if i, err := strconv.Atoi(tmp); err == nil {
				if i < len(c.List) {
					return c.List[i].ID
				}
			}
		}
		if s.Err() != nil {
			log.Fatal(s.Err())
		}
	}
}
