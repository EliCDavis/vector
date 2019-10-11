package vector

import "math"

type Vector2 struct {
	x float64
	y float64
}

func NewVector2(x float64, y float64) Vector2 {
	return Vector2{
		x: x,
		y: y,
	}
}

func (v Vector2) X() float64 {
	return v.x
}

func (v Vector2) Y() float64 {
	return v.y
}

// Add returns a vector that is the result of two vectors added together
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v Vector2) Length() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y))
}

func (v Vector2) Normalized() Vector2 {
	if v.Length() == 0 {
		return NewVector2(v.x, v.y)
	}
	return v.DivByConstant(v.Length())
}

func (v Vector2) MultByConstant(t float64) Vector2 {
	return Vector2{
		x: v.x * t,
		y: v.y * t,
	}
}

func (v Vector2) DivByConstant(t float64) Vector2 {
	return v.MultByConstant(1.0 / t)
}

// Distance is the euclidian distance between two points
func (v Vector2) Distance(other Vector2) float64 {
	return math.Sqrt(math.Pow(other.x-v.x, 2.0) + math.Pow(other.y-v.y, 2.0))
}
