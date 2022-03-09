package spritesheet

import (
	"image/color"
	"spritely/internal/shared/message"
	"spritely/internal/shared/topic"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
)

func (ss *SpriteSheet) mailbox() {
	msg := ss.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.LEFT_CLICK:
			ss.handleClick(m.GetPayload().(geom.Coordinate))
		case topic.SET_CURRENT_COLOR:
			ss.currentColor = m.GetPayload().(color.Color)
		case topic.SET_PIXEL:
			coord := m.GetPayload().(geom.Coordinate)
			ss.sprites[ss.selected.Y][ss.selected.X].Widget.Elements[coord.Y][coord.X].SetGraphic(ss.currentColor)
		case topic.COPY:
			ss.broker.Publish(message.Message{
				Topic:   topic.PUSH_TO_CLIPBOARD,
				Payload: ss.sprites[ss.selected.Y][ss.selected.X].Widget.Elements,
			})
		case topic.UPDATE_CANVAS:
			ss.sprites[ss.selected.Y][ss.selected.X].Widget.SetElements(m.GetPayload().([][]*widget.Element))
		case topic.LEFT:
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X - 1,
				Y: ss.selected.Y,
			})
		case topic.RIGHT:
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X + 1,
				Y: ss.selected.Y,
			})
		case topic.UP:
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X,
				Y: ss.selected.Y - 1,
			})
		case topic.DOWN:
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X,
				Y: ss.selected.Y + 1,
			})
		}
	}
}
