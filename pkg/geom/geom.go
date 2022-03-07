package geom

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type Bounds struct {
	Higher Coordinate
	Lower  Coordinate
}

type Offset struct {
	X float64
	Y float64
}

func (o Offset) String() string {
	return fmt.Sprintf("(%d, %d)", int(o.X), int(o.Y))
}

func (o Offset) ToCoordinate() Coordinate {
	return Coordinate{
		X: int(o.X),
		Y: int(o.Y),
	}
}

func ToCoordinate(x int, y int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}
