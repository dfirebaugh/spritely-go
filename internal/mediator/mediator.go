package mediator

import (
	"spritely/config"
	"spritely/pkg/broker"
	"spritely/pkg/widg"
	"spritely/pkg/widget"
	"spritely/pkg/widgetz"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mediator struct {
	broker  *broker.Broker
	widgets []*widget.Widget
}

func New() Mediator {
	b := broker.NewBroker()
	mediator := Mediator{
		broker: b,
	}
	layout := widg.Load("./config/layout.yml")
	println(layout.String())

	cfg := make(widg.Layout)
	cfg.Unmarshal(config.LayoutRaw)
	println(cfg.String())

	for _, v := range layout {
		mediator.widgets = append(mediator.widgets, widg.WidgetFromConfig(v))
	}

	go mediator.broker.Start()

	widgetz.New()
	return mediator
}

func (m *Mediator) Update() {
	for _, w := range m.widgets {
		w.Update()
	}
}

func (m *Mediator) Render(dst *ebiten.Image) {
	for _, w := range m.widgets {
		w.Render(dst)
	}
}
