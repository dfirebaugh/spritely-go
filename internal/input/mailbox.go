package input

import (
	"spritely/internal/topic"
	"spritely/pkg/widget"
)

func (i *Controller) mailbox() {
	msg := i.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.PUSH_TO_CLIPBOARD.String():
			i.clipboard.ReceivePixels(m.GetPayload().([][]*widget.Element))
		}
	}
}
