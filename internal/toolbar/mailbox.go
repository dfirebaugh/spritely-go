package toolbar

import (
	"github.com/dfirebaugh/spritely-go/internal/tool"
	"github.com/dfirebaugh/spritely-go/internal/topic"
	"github.com/dfirebaugh/spritely-go/pkg/geom"
)

func (tb *ToolBar) mailbox() {
	msg := tb.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.LEFT_CLICK.String():
			tb.handleClick(m.GetPayload().(geom.Coordinate))
		case topic.SET_CURRENT_TOOL.String():
			t := m.GetPayload().(tool.Tool)
			tb.pick(t)
			tb.Widget.SelectElement(geom.Coordinate{
				X: int(t),
				Y: 0,
			})

		}
	}
}
