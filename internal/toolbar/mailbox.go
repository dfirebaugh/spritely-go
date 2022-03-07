package toolbar

import (
	"spritely/internal/shared/topic"
	"spritely/internal/tool"
	"spritely/pkg/actor"
	"spritely/pkg/geom"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func (t *ToolBar) Message(msg actor.Message) {
	switch msg.Topic {
	case topic.RENDER:
		t.actorSystem.Lookup(t.widgetAddress).Message(actor.Message{
			Topic:   topic.RENDER,
			Payload: msg.Payload.(*ebiten.Image),
		})
	case topic.UPDATE:
		t.update()
	case topic.SET_CURRENT_TOOL:
		t.pick(msg.Payload.(tool.Tool))
	case topic.SET_OFFSET:
		t.actorSystem.Lookup(t.widgetAddress).Message(msg)
	case topic.HANDLE_CLICK:
		t.pick(tool.Tool(msg.Payload.(geom.Coordinate).X))
	}
}
func (t *ToolBar) SetAddress(address actor.Address) {
	t.address = address
}
