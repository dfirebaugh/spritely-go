package geom

import "testing"

func TestToCoordinate(t *testing.T) {
	offset := Offset{
		X: 1.1,
		Y: 4.2,
	}

	if offset.ToCoordinate().X != 1 && offset.ToCoordinate().Y != 4 {
		t.Errorf("should convert to int coordinate")
	}
}
