package quadtree

import (
	"fmt"
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
		{15, 60, 0},
		{19, 31, 1},
		{44, 47, 11},
		{0, 63, 0},
		{0, 127, 0},
	}
	for _, c := range cases {
		got := cellNum(c.UL, c.LR)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

var cellMaxs = [7]int{0, 1, 5, 21, 85, 341, 1365}

func cellNum(upperLeft, lowerRight int) int {
	n := upperLeft ^ lowerRight
	// fmt.Println(n)
	lv := msb(n) / 2
	fmt.Printf("Lv: %d\n", lv)
	shift := (lv + 1) * 2
	fmt.Printf("shift: %d\n", shift)
	fmt.Println(lowerRight >> uint(shift))

	switch {
	case n>>8&0x3 != 0:
		return 0
	case n>>6&0x3 != 0:
		return lowerRight>>8 + 1
	case n>>4&0x3 != 0:
		return lowerRight>>6 + 5
	case n>>2&0x3 != 0:
		return lowerRight>>4 + 21
	case n&0x3 != 0:
		return lowerRight>>2 + 85
	default:
		return lowerRight + 341
	}
}

func count(n int) int {
	n = n&0x5555 + n>>1&0x5555
	n = n&0x3333 + n>>2&0x3333
	n = n&0x0f0f + n>>4&0x0f0f
	return n&0x00ff + n>>8&0x00ff
}

func msb(n int) int {
	if n == 0 {
		return -1
	}
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	return count(n) - 1
}

func separate(n int) int {
	n = (n | (n << 8)) & 0x00ff00ff
	n = (n | (n << 4)) & 0x0f0f0f0f
	n = (n | (n << 2)) & 0x33333333
	n = (n | (n << 1)) & 0x55555555
	return n
}

func morton(x, y int) int {
	return separate(x) | separate(y)<<1
}
