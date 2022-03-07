package widget

import (
	"image/color"
	"spritely/pkg/draw"
	"spritely/pkg/geom"

	"github.com/hajimehoshi/ebiten/v2"
)

type Size struct {
	Width  int
	Height int
}

type Widget struct {
	Elements    [][]*Element
	elementSize Size
	selected    geom.Coordinate
	offset      geom.Offset
	selectable  bool
}

func NewSelectableColors(colors [][]color.Color, elementSize Size) *Widget {
	widget := NewWithColorElements(colors, elementSize)
	widget.selectable = true
	return widget
}
func NewSelectableImages(images [][]*ebiten.Image, elementSize Size) *Widget {
	widget := NewWithImageElements(images, elementSize)
	widget.selectable = true
	return widget
}
func NewWithImageElements(images [][]*ebiten.Image, elementSize Size) *Widget {
	widget := Widget{
		elementSize: elementSize,
		selectable:  false,
	}
	widget.initImages(images, elementSize)
	return &widget
}
func NewWithColorElements(colors [][]color.Color, elementSize Size) *Widget {
	widget := Widget{
		elementSize: elementSize,
		selectable:  false,
	}
	widget.initColors(colors, elementSize)
	return &widget
}

func (w *Widget) initColors(graphics [][]color.Color, elementSize Size) {
	var elements [][]*Element
	for _, v := range graphics {
		var elementRow []*Element
		for _, graphic := range v {
			elementRow = append(elementRow, &Element{
				size:    elementSize,
				Graphic: graphic,
			})
		}
		elements = append(elements, elementRow)
	}
	w.Elements = elements
}
func (w *Widget) initImages(graphics [][]*ebiten.Image, elementSize Size) {
	var elements [][]*Element
	for _, v := range graphics {
		var elementRow []*Element
		for _, graphic := range v {
			elementRow = append(elementRow, &Element{
				size:    elementSize,
				Graphic: graphic,
			})
		}
		elements = append(elements, elementRow)
	}
	w.Elements = elements
}

func (w *Widget) Render(dst *ebiten.Image) {
	for y, row := range w.Elements {
		for x, e := range row {
			e.render(dst, x, y)

			if !w.selectable || y != w.selected.Y || x != w.selected.X {
				continue
			}

			draw.DrawBox(
				dst,
				float64(x*w.elementSize.Width)+w.offset.X,
				float64(y*w.elementSize.Height)+w.offset.Y,
				float64(w.elementSize.Width),
				float64(w.elementSize.Height),
				color.White,
			)
		}
	}
}

func (w *Widget) SetOffset(offset geom.Offset) {
	w.offset = geom.Offset{
		X: offset.X,
		Y: offset.Y,
	}
	for y, row := range w.Elements {
		for x, e := range row {
			e.size = w.elementSize
			e.offset.X = w.offset.X + float64(x*e.size.Width)
			e.offset.Y = w.offset.Y + float64(y*e.size.Height)
		}
	}
}

func (w *Widget) DeriveBounds() geom.Bounds {
	return geom.Bounds{
		Lower: geom.Coordinate{
			X: w.offset.ToCoordinate().X,
			Y: w.offset.ToCoordinate().Y,
		},
		Higher: geom.Coordinate{
			X: w.offset.ToCoordinate().X + (len(w.Elements[0]) * w.elementSize.Width) - 1,
			Y: w.offset.ToCoordinate().Y + (len(w.Elements) * w.elementSize.Height) - 1,
		},
	}
}

func (w *Widget) IsWithinBounds(coordinate geom.Coordinate) bool {
	bounds := w.DeriveBounds()
	if coordinate.X < (bounds.Lower.X) || coordinate.X > (bounds.Higher.X) {
		return false
	}
	if coordinate.Y < (bounds.Lower.Y) || coordinate.Y > (bounds.Higher.Y) {
		return false
	}

	return true
}

func (w *Widget) SelectElement(coordinate geom.Coordinate) {
	w.selected = coordinate
}

func (w Widget) ToLocalCoordinate(coordinate geom.Coordinate) geom.Coordinate {
	return geom.Coordinate{
		X: (coordinate.X - w.offset.ToCoordinate().X) / w.elementSize.Width,
		Y: (coordinate.Y - w.offset.ToCoordinate().Y) / w.elementSize.Height,
	}
}

func ToCoordinate(x int, y int) geom.Coordinate {
	return geom.Coordinate{
		X: x,
		Y: y,
	}
}
