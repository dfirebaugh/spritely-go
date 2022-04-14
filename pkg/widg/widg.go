package widg

import (
	"fmt"
	"io/ioutil"
	"spritely/assets/icons"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/yaml.v3"
)

type Mailbox struct {
	Subscribe []string
	Publish   []string
}

type Offset struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}

type Widget struct {
	Label         string   `yaml:"label"`
	Subscriptions []string `yaml:"subsribe"`
	Mailbox       Mailbox  `yaml:"mailbox"`
	Offset        Offset   `yaml:"offset"`
	ElementSize   int      `yaml:"element-size"`
	Behaviors     []string `yaml:"behaviors"`
}

func (w Widget) String() string {
	str := fmt.Sprintf("label: %s\n", w.Label)
	str = str + fmt.Sprintf("element-size: %d\n", w.ElementSize)
	str = str + fmt.Sprintf("offset: %d, %d\n", w.Offset.X, w.Offset.Y)
	str = str + "subscriptions: \n"
	for _, s := range w.Mailbox.Subscribe {
		str = str + fmt.Sprintf("\t%s\n", s)
	}

	str = str + "publish: \n"
	for _, p := range w.Mailbox.Publish {
		str = str + fmt.Sprintf("\t%s\n", p)
	}
	str = str + "behaviors:\n"
	for _, b := range w.Behaviors {
		str = str + fmt.Sprintf("\t%s\n", b)
	}

	return str + "\n\n"
}

func LoadWidget(path string) *widget.Widget {
	y, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return nil
	}
	var w Widget

	yaml.Unmarshal(y, &w)
	println()
	println(w.String())
	println()

	img := icons.ToImage(w.Label)
	width, height := img.Size()

	return widget.NewSelectableImages(
		[][]*ebiten.Image{{img}},
		geom.Size{
			Width:  width,
			Height: height,
		},
		geom.Offset{
			X: float64(w.Offset.X),
			Y: float64(w.Offset.Y),
		})
}

func WidgetFromConfig(el *Element) *widget.Widget {
	return widget.NewSelectableImages(
		[][]*ebiten.Image{{el.Image}},
		geom.Size{
			Width:  el.ElementSize,
			Height: el.ElementSize,
		},
		geom.Offset{
			X: float64(el.Offset.X),
			Y: float64(el.Offset.Y),
		})
}
