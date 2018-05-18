package quadtree

import (
	"testing"
)

func TestSeparate(t *testing.T) {
	cases := []struct {
		Put, Want int
	}{
		{0x3, 0x5},
		{0xf, 0x55},
		{0xff, 0x5555},
		{0xfff, 0x555555},
		{0xffff, 0x55555555},
	}
	for _, c := range cases {
		got := separate(c.Put)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestMorton(t *testing.T) {
	cases := []struct {
		X, Y, Want int
	}{
		{3, 6, 45},
	}
	for _, c := range cases {
		got := morton(c.X, c.Y)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		Put, Want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{0xaa, 4},
		{0xf0, 4},
		{0xff, 8},
		{0xfff, 12},
		{0xffff, 16},
	}
	for _, c := range cases {
		got := count(c.Put)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestMsb(t *testing.T) {
	cases := []struct {
		Put, Want int
	}{
		{0, -1},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{8, 3},
		{16, 4},
		{17, 4},
		{32, 5},
		{64, 6},
		{128, 7},
		{0xff, 7},
		{0xfff, 11},
		{0xffff, 15},
	}
	for _, c := range cases {
		got := msb(c.Put)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestCellNum(t *testing.T) {
	cases := []struct {
		UL, LR, Want int
	}{
		{15, 60, 5},
		{19, 31, 22},
		{44, 47, 96},
		{0, 63, 5},
		{0, 0, 341},
		{0, 255, 1},
		{0, 1023, 0},
		{1023, 1023, 1364},
	}
	for _, c := range cases {
		got := cellNum(c.UL, c.LR)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
