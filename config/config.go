package config

import (
	_ "embed"
)

// embed configs so that they are packaged with the binary

//go:embed layout.yml
var LayoutRaw []byte

//go:embed colorpicker.yml
var ColorPickerRaw []byte
