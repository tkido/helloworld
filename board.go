package main

import "math/rand"

// Board in matrix of cell
type Board [][]int

// NewBoard get Board
func NewBoard() Board {
	b := make([][]int, h, h)
	for y := range b {
		b[y] = make([]int, w, w)
	}
	return b
}

// Clear all cells
func (b Board) Clear() {
	for y := range b {
		for x := range b[y] {
			b[y][x] = dead
		}
	}
}

// Rand randomize all cells
func (b Board) Rand() {
	for y := range b {
		for x := range b[y] {
			if rand.Float64() < 0.3 {
				b[y][x] = live
			} else {
				b[y][x] = dead
			}
		}
	}
}
