package geom

import "fmt"

type Geometry struct {
	Size   Size
	Offset Offset
}
type Coordinate struct {
	X int
	Y int
}

type Size struct {
	Width  int
	Height int
}

type Bounds struct {
	Higher Coordinate
	Lower  Coordinate
}

type Offset struct {
	X float64
	Y float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
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
