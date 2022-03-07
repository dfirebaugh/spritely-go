package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawBox(dst *ebiten.Image, x, y, width, height float64, clr color.Color) {
	var stroke float64 = 1
	// top
	ebitenutil.DrawLine(
		dst,
		x,
		y,
		x+width,
		y,
		clr,
	)
	// bottom
	ebitenutil.DrawLine(
		dst,
		x,
		y+height-stroke,
		x+width,
		y+height-stroke,
		clr,
	)
	// left
	ebitenutil.DrawLine(
		dst,
		x+stroke,
		y,
		x+stroke,
		y+height,
		clr,
	)
	// right
	ebitenutil.DrawLine(
		dst,
		x+width,
		y,
		x+width,
		y+height,
		clr,
	)
}
