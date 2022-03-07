package command

import "spritely/pkg/actor"

type Command struct{}
type CommandStack struct {
	actorSystem *actor.ActorSystem
	commands    []Command
}

func New(actorSystem *actor.ActorSystem) CommandStack {
	return CommandStack{actorSystem: actorSystem}
}

func (c *CommandStack) push()    {}
func (c *CommandStack) pop()     {}
func (c *CommandStack) isEmpty() {}
