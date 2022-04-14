package widgets

import (
	"spritely/pkg/widg"
	"spritely/pkg/widget"
)

func LoadWidgets() []*widget.Widget {
	var widgets []*widget.Widget
	w := widg.LoadWidget("./config/toolbar.widget.yml")
	p := widg.LoadWidget("./config/palette.widget.yml")

	widgets = append(widgets, w)
	widgets = append(widgets, p)

	return widgets
}
