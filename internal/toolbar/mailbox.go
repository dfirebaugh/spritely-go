package toolbar

import (
	"spritely/internal/shared/topic"
	"spritely/pkg/geom"
)

func (t *ToolBar) mailbox() {
	msg := t.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.LEFT_CLICK:
			t.handleClick(m.GetPayload().(geom.Coordinate))
		}
	}
}
