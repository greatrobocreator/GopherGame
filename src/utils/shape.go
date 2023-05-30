package utils

//package utils

type Shape interface{}

type Rectangle struct {
	Pos, Size Vector
}

func NewRectangle(pos, size Vector) Rectangle {
	return Rectangle{Pos: pos, Size: size}
}

func RectanglesIntersection(l Rectangle, r Rectangle) bool {

	pointInSegment := func(s Vector, p float64) bool {
		return (p-s.X)*(s.Y-p) >= 0
	}

	segmentsIntersection := func(l, r Vector) bool {
		return pointInSegment(r, l.X) || pointInSegment(r, l.Y)
	}

	lPoint := l.Pos.Add(l.Size)
	rPoint := r.Pos.Add(r.Size)
	return segmentsIntersection(NewVector(l.Pos.X, lPoint.X), NewVector(r.Pos.X, rPoint.X)) &&
		segmentsIntersection(NewVector(l.Pos.Y, lPoint.Y), NewVector(r.Pos.Y, rPoint.Y))
}
