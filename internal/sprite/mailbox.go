package sprite

import (
	"image/color"
	"spritely/internal/shared/topic"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
)

var lastMsg string

func (s *Sprite) mailbox() {
	msg := s.broker.Subscribe()
	for {
		m := <-msg
		if lastMsg == m.Hash() {
			return
		}
		switch m.GetTopic() {
		case topic.LEFT_CLICK:
			s.handleClick(m.GetPayload().(geom.Coordinate))
		case topic.RIGHT_CLICK:
			s.handleRightClick(m.GetPayload().(geom.Coordinate))
		case topic.SET_CURRENT_COLOR:
			s.currentColor = m.GetPayload().(color.Color)
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
