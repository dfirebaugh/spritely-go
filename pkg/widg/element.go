package widg

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"spritely/assets/icons"
	"text/template"

	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/yaml.v3"
)

type Element struct {
	Name   string
	Offset struct {
		X int `yaml:"x"`
		Y int `yaml:"y"`
	} `yaml:"offset"`
	Children    []Element `yaml:"children"`
	ImagePath   string    `yaml:"img"`
	Image       *ebiten.Image
	ElementSize int `yaml:"element-size"`
	Mailbox     struct {
		Publish   []string `yaml:"publish"`
		Subscribe []string `yaml:"subscribe"`
	} `yaml:"mailbox"`
	Behaviors []string `yaml:"behaviors"`
}

func (e Element) String() string {
	tmpl := template.New("")
	text := `{{.Name}}:
	offset: {{.Offset.X}}, {{.Offset.Y}}
	image: {{.Image}}
	mail-box:
		Publishes:
			{{range $index, $element := .Mailbox.Publish}} 
			{{$element}}{{end}}
		Subscribes:
			{{range $index, $element := .Mailbox.Subscribe}}
			{{$element}}{{end}}
	behaviors:
		{{range $index, $element := .Behaviors}}
			{{$element}}{{end}}
`
	tmpl.Parse(text)
	var b bytes.Buffer
	tmpl.Execute(io.Writer(&b), e)

	return b.String()
}

func (e *Element) Unmarshal(raw []byte) {
	yaml.Unmarshal(raw, e)
	e.Image = icons.ToImage(e.ImagePath)
}

// look at the label of the Element and verify that
//   there is a file for that Element
//   we look for <element-name>.element.yml
//
func (e *Element) loadElement() {
	raw, err := ioutil.ReadFile(fmt.Sprintf("./config/%s.yml", e.Name))
	if err != nil {
		println(err.Error())
		return
	}
	e.Unmarshal(raw)
}
