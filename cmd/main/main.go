package main

import (
	_ "image/png"
	"spritely/internal/game"
)

const (
	ScreenWidth  = 1024
	ScreenHeight = 1024
)

// Tile is 8x8 pixels
// MettaSpriteSheet
// Pattern Table is an optimized version of a Meta Sprite Sheet
// Pallette is a set of 4 colors - The zeroth index is always interpretted as transparent

// Allow user to draw in a natural way and have those assets get optimized down into a pattern table

func main() {
	game.New(ScreenWidth, ScreenHeight).Run()
}
