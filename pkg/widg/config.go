package widg

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func loadFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("%s\n", err)
		return nil
	}
	return file
}

func Load(path string) Layout {
	cfg := make(Layout)
	file := loadFile(path)

	cfg.Unmarshal(file)

	return cfg
}
