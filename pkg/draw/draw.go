package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawBox(dst *ebiten.Image, x, y, width, height float64, clr color.Color) {
	var stroke float64 = 1
	// top
	vector.StrokeLine(dst, float32(x), float32(y), float32(x+width), float32(y), float32(stroke), clr, false)
	// bottom
	vector.StrokeLine(dst, float32(x), float32(y+height-stroke+1), float32(x+width), float32(y+height-stroke), float32(stroke), clr, false)
	// left
	vector.StrokeLine(dst, float32(x+stroke), float32(y), float32(x+stroke), float32(y+height), float32(stroke), clr, false)
	// right
	vector.StrokeLine(dst, float32(x+width), float32(y), float32(x+width), float32(y+height), float32(stroke), clr, false)
}
