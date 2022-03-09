package widget

import (
	"fmt"
	"image/color"
	"spritely/pkg/draw"
	"spritely/pkg/geom"

	"github.com/hajimehoshi/ebiten/v2"
)

type Widget struct {
	Elements    [][]*Element
	elementSize geom.Size
	selected    geom.Coordinate
	offset      geom.Offset
	selectable  bool
}

func NewSelectableColors(colors [][]color.Color, elementSize geom.Size, offset geom.Offset) *Widget {
	widget := NewWithColorElements(colors, elementSize, offset)
	widget.selectable = true
	return widget
}
func NewSelectableImages(images [][]*ebiten.Image, elementSize geom.Size, offset geom.Offset) *Widget {
	widget := NewWithImageElements(images, elementSize, offset)
	widget.selectable = true
	return widget
}
func NewWithImageElements(images [][]*ebiten.Image, elementSize geom.Size, offset geom.Offset) *Widget {
	widget := Widget{
		elementSize: elementSize,
		selectable:  false,
	}
	widget.initImages(images, elementSize)
	widget.SetOffset(offset)
	return &widget
}
func NewWithColorElements(colors [][]color.Color, elementSize geom.Size, offset geom.Offset) *Widget {
	widget := Widget{
		elementSize: elementSize,
		selectable:  false,
	}
	widget.initColors(colors, elementSize)
	widget.SetOffset(offset)
	return &widget
}

func (w *Widget) initColors(graphics [][]color.Color, elementSize geom.Size) {
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
func (w *Widget) initImages(graphics [][]*ebiten.Image, elementSize geom.Size) {
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

func (w *Widget) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		coord := geom.Coordinate{
			X: x,
			Y: y,
		}

		if !w.IsWithinBounds(coord) {
			return
		}
		local := w.ToLocalCoordinate(coord)
		w.SelectElement(local)
	}
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

func (w *Widget) DeriveLocalBounds() geom.Bounds {
	return geom.Bounds{
		Lower: geom.Coordinate{
			X: 0,
			Y: 0,
		},
		Higher: geom.Coordinate{
			X: len(w.Elements[0]) - 1,
			Y: len(w.Elements) - 1,
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

func (w *Widget) IsWithinLocalBounds(coordinate geom.Coordinate) bool {
	bounds := w.DeriveLocalBounds()

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

func (w *Widget) SetElements(elements [][]*Element) {
	// w.Elements = w.DeepCopy(elements)
	for y, row := range elements {
		for x, p := range row {
			w.Elements[y][x].SetGraphic(p.Graphic)
		}
	}
}

func (w *Widget) GetElement(local geom.Coordinate) *Element {
	return w.Elements[local.Y][local.X]
}

func (w *Widget) DeepCopy(elements [][]*Element) [][]*Element {
	var elms [][]*Element
	for _, row := range elements {
		var elemRow []*Element
		for _, p := range row {
			elemRow = append(elemRow, &Element{
				Graphic: p.Graphic,
			})
		}
		elms = append(elms, elemRow)
	}
	return elms
}

func (w *Widget) DebugPrint() {
	for _, row := range w.Elements {
		for _, p := range row {
			r, g, b, _ := p.Graphic.(color.Color).RGBA()
			fmt.Printf("%d,%d,%d ", r, g, b)
		}
		println()
	}
	println()
}

func DebugPrint(elements [][]*Element) {
	for _, row := range elements {
		for _, p := range row {
			r, g, b, _ := p.Graphic.(color.Color).RGBA()
			fmt.Printf("%d,%d,%d ", r, g, b)
		}
		println()
	}
	println()
}

func (w Widget) GetSize() int {
	return len(w.Elements) * len(w.Elements[0])
}
