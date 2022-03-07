package widgetmediator

import (
	"spritely/internal/shared/request"
	"spritely/internal/shared/topic"
	"spritely/internal/tool"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

func (w *WidgetMediator) Message(msg actor.Message) {
	switch msg.Topic {
	case topic.RENDER:
		w.widget.Render(msg.Payload.(*ebiten.Image))
	case topic.SET_OFFSET:
		w.widget.SetOffset(msg.Payload.(geom.Offset))
	case topic.SET_CURRENT_TOOL:
		// w.widget.SelectElement(geom.Coordinate{
		// 	X: int(msg.Payload.(tool.Tool)),
		// 	Y: 0,
		// })

		println(int(msg.Payload.(tool.Tool)))
	case topic.HANDLE_CLICK:
		w.handleClick(msg.Payload.(geom.Coordinate), msg.Requestor)
	case topic.GET_PIXELS:
		w.actorSystem.Lookup(msg.Requestor).Message(actor.Message{
			Topic:   topic.PUSH_PIXELS,
			Payload: w.widget.Elements,
		})
	case topic.PUSH_PIXELS:
		newElements := msg.Payload.([][]*widget.Element)
		for y, row := range w.widget.Elements {
			for x, element := range row {
				element.SetGraphic(newElements[y][x].Graphic)
			}
		}
	case topic.SET_PIXEL:
		req := msg.Payload.(request.SetPixel)
		w.widget.Elements[req.Coordinate.Y][req.Coordinate.X].SetGraphic(req.Color)
	}
}

func (w *WidgetMediator) SetAddress(address actor.Address) {
	w.address = address
}
