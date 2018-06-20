package ui

import (
	"image/color"
	"log"
	"strconv"
)

var (
	colorCache    map[string]color.Color
	colorCodeRune map[rune]struct{}
)

func init() {
	colorCache = map[string]color.Color{}
	colorCodeRune = map[rune]struct{}{}
	for _, r := range "0123456789abcdefABCDEF" {
		colorCodeRune[r] = struct{}{}
	}

}

// Color return color
func Color(s string) color.Color {
	raw := s
	if c, ok := colorCache[s]; ok {
		return c
	}
	for _, r := range s {
		if _, ok := colorCodeRune[r]; !ok {
			log.Fatalf("invalid rune %s for color code", string(r))
		}
	}
	switch len(s) {
	case 3, 4:
		rs := []rune{}
		for _, r := range s {
			rs = append(rs, r, r)
		}
		s = string(rs)
	case 6, 8:
	default:
		log.Fatalf("invalid color code %s", s)
	}
	us := []uint8{}
	for i := 0; i < len(s)/2; i++ {
		ui, _ := strconv.ParseUint(s[i*2:i*2+2], 16, 8)
		us = append(us, uint8(ui))
	}
	if len(us) == 3 {
		us = append(us, 0xff)
	}
	c := color.NRGBA{us[0], us[1], us[2], us[3]}
	colorCache[raw] = c
	return c
}
