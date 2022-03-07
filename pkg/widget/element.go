package widget

import (
	"image/color"
	"spritely/pkg/geom"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Element struct {
	Graphic interface{}
	size    Size
	offset  geom.Offset
}

func (e *Element) SetGraphic(graphic interface{}) {
	e.Graphic = graphic
}

func (e Element) render(dst *ebiten.Image, x int, y int) {
	// padding := 4
	image, ok := e.Graphic.(*ebiten.Image)
	if ok {
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(
			e.offset.X+float64(e.size.Width/4),
			e.offset.Y+float64(e.size.Height/4),
		)
		dst.DrawImage(image, op)
	}

	pixel, ok := e.Graphic.(color.Color)
	if ok {
		ebitenutil.DrawRect(
			dst,
			e.offset.X,
			e.offset.Y,
			float64(e.size.Width),
			float64(e.size.Height),
			pixel,
		)
	}
}
