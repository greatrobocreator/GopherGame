package utils

type Vector struct {
	X, Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector) MulVector(other Vector) Vector {
	return Vector{X: v.X * other.X, Y: v.Y * other.Y}
}

func (v Vector) Mul(x float64) Vector {
	return Vector{X: v.X * x, Y: v.Y * x}
}

func (v Vector) Flatten() [2]float64 {
	return [2]float64{v.X, v.Y}
}
