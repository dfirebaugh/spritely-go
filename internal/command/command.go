package command

import "github.com/dfirebaugh/spritely-go/pkg/actor"

type (
	Command      struct{}
	CommandStack struct {
		actorSystem *actor.ActorSystem
		commands    []Command
	}
)

func New(actorSystem *actor.ActorSystem) CommandStack {
	return CommandStack{actorSystem: actorSystem}
}

func (c *CommandStack) push()    {}
func (c *CommandStack) pop()     {}
func (c *CommandStack) isEmpty() {}
