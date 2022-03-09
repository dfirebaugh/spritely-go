package clipboard

import (
	"spritely/pkg/widget"
)

type Controller struct {
	Pixels [][]*widget.Element
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) ReceivePixels(elements [][]*widget.Element) {
	var elms [][]*widget.Element
	for _, row := range elements {
		var elemRow []*widget.Element
		for _, p := range row {
			elemRow = append(elemRow, &widget.Element{
				Graphic: p.Graphic,
			})
		}
		elms = append(elms, elemRow)
	}
	c.Pixels = elms
}
