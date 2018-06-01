package ui

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
)

// ColorCode show RGBA color HTML color codes like.
// e.g. White -> #ffffffff
func ColorCode(c color.Color) string {
	rgba := [4]uint32{}
	rgba[0], rgba[1], rgba[2], rgba[3] = c.RGBA()
	buf := bytes.Buffer{}
	buf.WriteString("#")
	for _, u := range rgba {
		buf.WriteString(fmt.Sprintf("%02x", uint8(u)))
	}
	return buf.String()
}

// isCloseEnough returns that distance between given two points
// (button downed and button upped) is close enough
// to be recognized one click
func isCloseEnough(a, b image.Point) bool {
	sub := a.Sub(b)
	if sub.X*sub.X+sub.Y*sub.Y <= 16 {
		return true
	}
	return false
}
