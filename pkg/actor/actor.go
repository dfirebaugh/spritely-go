package actor

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Address uuid.UUID

type Actor interface {
	Message(Message)
	SetAddress(Address)
}

type ActorSystem struct {
	addressBook map[Address]Actor
}

type Topic int
type Mailbox chan Message

type Message struct {
	Requestor Address
	Topic     Topic
	Payload   interface{}
}

type messageStub struct{}

func (messageStub) SetAddress(address Address) {}
func (messageStub) Message(msg Message) {
	log.Errorf("unable to find actor: %s %#v", msg.Topic, msg.Payload)
}

var actorSystem *ActorSystem = nil

func New() *ActorSystem {
	if actorSystem == nil {
		return &ActorSystem{
			addressBook: make(map[Address]Actor),
		}
	}

	return actorSystem
}

func (a ActorSystem) Register(actor Actor) Address {
	address := Address(uuid.New())
	a.addressBook[address] = actor
	actor.SetAddress(address)
	return address
}

func (a ActorSystem) Lookup(uid Address) Actor {
	if val, ok := a.addressBook[uid]; ok {
		return val
	}

	return messageStub{}
}

func (a Address) String() string {
	return uuid.UUID(a).String()
}
