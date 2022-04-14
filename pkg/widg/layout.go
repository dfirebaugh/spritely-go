package widg

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Layout map[string]*Element

func (l Layout) String() string {
	var result string

	for _, v := range l {
		result = fmt.Sprintf("%s\n%s", result, v.String())
	}

	return result
}

func (l *Layout) Unmarshal(raw []byte) {
	yaml.Unmarshal(raw, l)

	for k, v := range *l {
		v.Name = k
		v.loadElement()
	}
}
