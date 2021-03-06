package types

import "github.com/Jaacky/typingwars/backend/pb"

// Point struct
type Point struct {
	X float32
	Y float32
}

// NewPoint initialization
func NewPoint(x, y float32) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (point *Point) Add(vector *Vector) *Point {
	return NewPoint(
		point.X+vector.X,
		point.Y+vector.Y,
	)
}

func (point *Point) Equal(other *Point) bool {
	return *point == *other
}

func (point *Point) GetXDirectionTo(other *Point) float32 {
	if point.X > other.X {
		return -1
	} else if point.X < other.X {
		return 1
	} else {
		return 1
	}
}

func (point *Point) ToProto() *pb.Point {
	return &pb.Point{
		X: point.X,
		Y: point.Y,
	}
}
