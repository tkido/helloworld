package ui

import (
	"image/color"
	"log"
	"testing"
)

func TestColor(t *testing.T) {
	cases := []struct {
		Put  string
		Want color.Color
	}{
		{"fff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"ffff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"ffffff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"ffffffff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"fff0", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"ffffff00", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"FFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFFFFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFFFFFFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFF0", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"FFFFFF00", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"abcdef", color.NRGBA{0xab, 0xcd, 0xef, 0xff}},
		{"ABCDEF", color.NRGBA{0xAB, 0xCD, 0xEF, 0xFF}},
	}
	for _, c := range cases {
		put := c.Put
		got, err := Color(c.Put)
		if err != nil {
			t.Error(err)
			continue
		}
		want := c.Want
		if got != want {
			t.Errorf("put %v got %v want %v", put, got, want)
		}
	}
}

func TestC(t *testing.T) {
	s := "0123456789abcdefABCDEF"
	for _, r := range s {
		log.Println(r)
	}

}
