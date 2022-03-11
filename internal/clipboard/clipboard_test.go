package clipboard

import (
	"spritely/pkg/actor"
	"testing"
)

type TestActor struct {
	addr    actor.Address
	lastMsg interface{}
}

func (t *TestActor) SetAddress(addr actor.Address) {
	t.addr = addr
}
func (t *TestActor) Message(msg actor.Message) {
	t.lastMsg = msg
}

func TestPushPixels(t *testing.T) {
	// as := actor.New()
	// cp := New(as)
	// ta := &TestActor{}
	// testActor := as.Register(ta)

	// pixels := [][]*widget.Element{{{Graphic: color.Black}}}

	// as.Lookup(cp).Message(actor.Message{
	// 	Topic:     topic.PUSH_PIXELS,
	// 	Requestor: testActor,
	// 	Payload:   pixels,
	// })

	// if len(ta.lastMsg.([][]*widget.Element)) <= 0 {
	// 	fmt.Printf("%#v", ta.lastMsg)
	// 	t.Errorf("did not receive the expected pixels")
	// 	return
	// }
}
