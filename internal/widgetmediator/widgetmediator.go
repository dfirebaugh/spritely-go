package widgetmediator

import (
	"image/color"
	"spritely/internal/shared/topic"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type WidgetMediator struct {
	actorSystem *actor.ActorSystem
	address     actor.Address
	widget      *widget.Widget
}

func NewSelectableColors(actorSystem *actor.ActorSystem, colors [][]color.Color, elementSize widget.Size) actor.Address {
	widget := WidgetMediator{
		actorSystem: actorSystem,
		widget:      widget.NewSelectableColors(colors, elementSize),
	}
	return actorSystem.Register(&widget)
}
func NewSelectableImages(actorSystem *actor.ActorSystem, images [][]*ebiten.Image, elementSize widget.Size) actor.Address {
	widget := WidgetMediator{
		actorSystem: actorSystem,
		widget:      widget.NewSelectableImages(images, elementSize),
	}
	return actorSystem.Register(&widget)
}
func NewWithImageElements(actorSystem *actor.ActorSystem, images [][]*ebiten.Image, elementSize widget.Size) actor.Address {
	widget := WidgetMediator{
		actorSystem: actorSystem,
		widget:      widget.NewWithImageElements(images, elementSize),
	}
	return actorSystem.Register(&widget)
}
func NewWithColorElements(actorSystem *actor.ActorSystem, colors [][]color.Color, elementSize widget.Size) actor.Address {
	widget := WidgetMediator{
		actorSystem: actorSystem,
		widget:      widget.NewWithColorElements(colors, elementSize),
	}
	return actorSystem.Register(&widget)
}

// handleClick - we verify that it is a click within the widget's local coordinates
//   if so, we send a message back to the requestor including
//   the local coordinate of the click
func (w *WidgetMediator) handleClick(coordinate geom.Coordinate, requestor actor.Address) {
	if !w.widget.IsWithinBounds(coordinate) {
		return
	}
	local := w.widget.ToLocalCoordinate(coordinate)
	w.actorSystem.Lookup(requestor).Message(actor.Message{
		Topic:   topic.HANDLE_CLICK,
		Payload: local,
	})
	w.widget.SelectElement(local)
}
