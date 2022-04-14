package widget

import (
	"image"
	"image/color"
	"spritely/pkg/geom"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Element struct {
	Graphic interface{}
	Size    geom.Size
	Offset  geom.Offset
}

func (e *Element) SetGraphic(graphic interface{}) {
	e.Graphic = graphic
}

func (e Element) Render(dst *ebiten.Image, x int, y int) {
	if e.Graphic == nil {
		return
	}

	if img, ok := e.Graphic.(*ebiten.Image); ok {
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(
			e.Offset.X+float64(e.Size.Width/4),
			e.Offset.Y+float64(e.Size.Height/4),
		)
		dst.DrawImage(img, op)

		dst.DrawImage(img.SubImage(image.Rect(0, 0, 10, 10)).(*ebiten.Image), op)
	}

	if pixel, ok := e.Graphic.(color.Color); ok {
		e.RenderPixel(dst, pixel)
	}
}

func (e Element) RenderPixel(dst *ebiten.Image, pixel color.Color) {
	ebitenutil.DrawRect(
		dst,
		e.Offset.X,
		e.Offset.Y,
		float64(e.Size.Width),
		float64(e.Size.Height),
		pixel,
	)
}

func (e Element) ColorMatches(e2 *Element) bool {
	return e.Graphic.(color.Color) == e2.Graphic.(color.Color)
}
