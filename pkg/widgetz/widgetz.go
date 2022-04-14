package widgetz

import "spritely/pkg/geom"

type Element struct {
	Graphic interface{}
	Size    geom.Size
	Offset  geom.Offset
}

type Widget struct {
	Elements    [][]*Element
	elementSize geom.Size
	selected    geom.Coordinate
	Offset      geom.Offset
	selectable  bool
}

func New() {

}

func (Widget) Update() {

}

func (Widget) Render() {
}
