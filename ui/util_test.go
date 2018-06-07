package ui

import (
	"image/color"
	"testing"
)

func TestColorCode(t *testing.T) {
	cases := []struct {
		Put  color.Color
		Want string
	}{
		{color.Black, "#000000ff"},
		{color.White, "#ffffffff"},
		{color.NRGBA{0xff, 0x00, 0x00, 0xff}, "#ff0000ff"},
		{color.NRGBA{0x00, 0xff, 0x00, 0xff}, "#00ff00ff"},
		{color.NRGBA{0x00, 0x00, 0xff, 0xff}, "#0000ffff"},
		{color.NRGBA{0x00, 0x00, 0x00, 0x00}, "#00000000"},
		{color.NRGBA{0xff, 0xff, 0xff, 0xff}, "#ffffffff"},
	}
	for _, c := range cases {
		put := c.Put
		got := ColorCode(c.Put)
		want := c.Want
		if got != want {
			t.Errorf("put %v got %v want %v", put, got, want)
		}
	}
}

func TestEqual(t *testing.T) {
	put := [3]byte{0, 1, 2}
	got := put
	want := [3]byte{0, 1, 2}
	if got != want {
		t.Errorf("put %v got %v want %v", put, got, want)
	}
}
