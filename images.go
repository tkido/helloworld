package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

func loadImage(path string) image.Image {
	f, err := os.Open(filepath.Join("images", path))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
