package sprite

import (
	"image/color"
	"spritely/internal/tool"
	"spritely/internal/topic"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
)

func (s *Sprite) mailbox() {
	msg := s.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.LEFT_CLICK.String():
			s.handleClick(m.GetPayload().(geom.Coordinate))
		case topic.RIGHT_CLICK.String():
			s.handleRightClick(m.GetPayload().(geom.Coordinate))
		case topic.SET_CURRENT_COLOR.String():
			s.currentColor = m.GetPayload().(color.Color)
		case topic.SET_CURRENT_TOOL.String():
			s.setCurrentTool(m.GetPayload().(tool.Tool))
		case topic.UPDATE_CANVAS.String():
			s.udpateCanvas(m.GetPayload().([][]*widget.Element))
		}
	}
}
