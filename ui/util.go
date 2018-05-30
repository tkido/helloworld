package ui

import (
	"bytes"
	"fmt"
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
