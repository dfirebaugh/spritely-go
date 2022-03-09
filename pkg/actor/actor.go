package actor

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Address uuid.UUID

// type Actor interface {
// 	Message(Message)
// 	SetAddress(Address)
// }

type Actor struct {
	address  Address
	children map[Address]Actor
}

func (a *Actor) SpawnChild() {

}
func (a *Actor) Publish() {

}
func (a *Actor) SubScribe() {

}

func (a *Actor) SetAddress(addr Address) {
	a.address = addr
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

func (m Message) Hash() string {
	h := sha256.New()
	h.Write([]byte(m.String()))
	return fmt.Sprintf("%x", h.Sum(nil))
}
func (m Message) String() string {
	return fmt.Sprintf("%#v", m)
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

// func (a ActorSystem) Lookup(uid Address) Actor {
// 	if val, ok := a.addressBook[uid]; ok {
// 		return val
// 	}

// 	return messageStub{}
// }

// String
// a string representation of the actor's address
func (a Address) String() string {
	return uuid.UUID(a).String()
}
