package request

import (
	"image/color"
	"spritely/pkg/geom"
)

type SetPixel struct {
	Coordinate    geom.Coordinate
	Color         color.Color
	CurrentSprite geom.Coordinate
}
