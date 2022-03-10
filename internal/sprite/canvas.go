package sprite

import (
	"image/color"
	"spritely/internal/message"
	"spritely/internal/palette"
	"spritely/internal/tool"
	"spritely/internal/topic"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
)

func NewCanvas(b *broker.Broker, offset geom.Offset, elementSize geom.Size) *Sprite {
	c := New(b, offset, elementSize)
	c.isCanvas = true
	c.setCurrentColor(palette.DefaultColors[0])
	return c
}

func (s *Sprite) executeCanvasOp(coord geom.Coordinate) {
	if !s.Widget.IsWithinBounds(coord) {
		return
	}
	if !s.isCanvas {
		return
	}
	local := s.Widget.ToLocalCoordinate(coord)
	switch s.currentTool {
	case tool.Pencil:
		s.pencilOp(local)
	case tool.Fill:
		s.fillOp(local)
	case tool.Drag:
		s.dragOp(local)
	}
}

func (s *Sprite) setCurrentColor(c color.Color) {
	s.currentColor = c
	go s.broker.Publish(message.Message{
		Topic:   topic.SET_CURRENT_COLOR,
		Payload: s.currentColor,
	})
}

func (s *Sprite) setCurrentTool(t tool.Tool) {
	s.currentTool = t
}

func (s *Sprite) pencilOp(local geom.Coordinate) {
	s.setPixel(local)
	s.broker.Publish(message.Message{
		Topic:   topic.SET_CURRENT_COLOR,
		Payload: s.Widget.Elements[local.Y][local.X].Graphic,
	})
}

var filledElements map[geom.Coordinate]*widget.Element

func (s *Sprite) fillOp(local geom.Coordinate) {
	filledElements = make(map[geom.Coordinate]*widget.Element, s.Widget.GetSize())
	s.fillNeighbors(local)
	s.broker.Publish(message.Message{
		Topic:   topic.UPDATE_CANVAS,
		Payload: s.Widget.Elements,
	})
}

// getNeighbors returns neighbors local coordinates of left, right, up, and down
func (s *Sprite) getNeighbors(local geom.Coordinate) []geom.Coordinate {
	return []geom.Coordinate{
		{
			X: local.X - 1,
			Y: local.Y,
		},
		{
			X: local.X + 1,
			Y: local.Y,
		},
		{
			X: local.X,
			Y: local.Y - 1,
		},
		{
			X: local.X,
			Y: local.Y - 1,
		},
		{
			X: local.X,
			Y: local.Y + 1,
		},
		// {
		// 	X: local.X + 1,
		// 	Y: local.Y + 1,
		// },
		// {
		// 	X: local.X - 1,
		// 	Y: local.Y - 1,
		// },
		// {
		// 	X: local.X - 1,
		// 	Y: local.Y + 1,
		// },
		// {
		// 	X: local.X + 1,
		// 	Y: local.Y - 1,
		// },
	}
}

func (s *Sprite) fillElement(local geom.Coordinate, originalElement *widget.Element) bool {
	if ok := s.Widget.IsWithinLocalBounds(local); !ok {
		return false
	}

	// the color of the pixel we are evaluating should be the same as the original element
	if ok := s.Widget.GetElement(local).ColorMatches(originalElement); !ok {
		return false
	}

	if _, ok := filledElements[local]; ok {
		return false
	}

	filledElements[local] = s.Widget.GetElement(local)
	return true
}

func (s *Sprite) fillNeighbors(local geom.Coordinate) {
	if !s.Widget.IsWithinLocalBounds(local) {
		return
	}
	e := s.Widget.GetElement(local)

	for _, l := range s.getNeighbors(local) {
		if ok := s.fillElement(l, e); !ok {
			continue
		}
		s.fillNeighbors(l)
	}

	e.SetGraphic(s.currentColor)
}

func (s *Sprite) dragOp(local geom.Coordinate) {
}

func (s *Sprite) udpateCanvas(elements [][]*widget.Element) {
	if !s.isCanvas {
		return
	}
	s.Widget.SetElements(elements)
}
