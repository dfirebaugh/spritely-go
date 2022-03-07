package input

import (
	"spritely/internal/shared/topic"
	"spritely/pkg/actor"
)

func (i *InputController) Message(msg actor.Message) {
	switch msg.Topic {
	case topic.UPDATE:
		i.Update()
	case topic.PUSH_PIXELS:
		i.actorSystem.Lookup(i.clipboard).Message(msg)
	case topic.GET_PIXELS:
		i.actorSystem.Lookup(i.clipboard).Message(msg)
	}
}
func (i *InputController) SetAddress(address actor.Address) {
	i.address = address
}
