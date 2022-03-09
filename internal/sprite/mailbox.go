package sprite

import (
	"image/color"
	"spritely/internal/shared/topic"
	"spritely/internal/tool"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
)

func (s *Sprite) mailbox() {
	msg := s.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.LEFT_CLICK:
			s.handleClick(m.GetPayload().(geom.Coordinate))
		case topic.RIGHT_CLICK:
			s.handleRightClick(m.GetPayload().(geom.Coordinate))
		case topic.SET_CURRENT_COLOR:
			s.currentColor = m.GetPayload().(color.Color)
		case topic.SET_CURRENT_TOOL:
			s.currentTool = m.GetPayload().(tool.Tool)
		case topic.PASTE:
			if !s.isCanvas {
				return
			}
			s.Widget.SetElements(m.GetPayload().([][]*widget.Element))
		case topic.SET_CANVAS:
			if !s.isCanvas {
				return
			}
			s.Widget.SetElements(m.GetPayload().([][]*widget.Element))
		}
	}
}
