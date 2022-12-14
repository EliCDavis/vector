package vector

import (
	"math"
	"math/rand"
)

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

func Vector2Zero() Vector2 {
	return Vector2{
		x: 0,
		y: 0,
	}
}

func Vector2Up() Vector2 {
	return Vector2{
		x: 0,
		y: 1,
	}
}

func Vector2Down() Vector2 {
	return Vector2{
		x: 0,
		y: -1,
	}
}

func Vector2Left() Vector2 {
	return Vector2{
		x: -1,
		y: 0,
	}
}

func Vector2Right() Vector2 {
	return Vector2{
		x: 1,
		y: 0,
	}
}

func Vector2One() Vector2 {
	return Vector2{
		x: 1,
		y: 1,
	}
}

func Vector2Rnd() Vector2 {
	return Vector2{
		x: rand.Float64(),
		y: rand.Float64(),
	}
}

func (v Vector2) X() float64 {
	return v.x
}

// SetX changes the x component of the vector
func (v Vector2) SetX(newX float64) Vector2 {
	return Vector2{
		x: newX,
		y: v.y,
	}
}

func (v Vector2) Y() float64 {
	return v.y
}

// SetY changes the y component of the vector
func (v Vector2) SetY(newY float64) Vector2 {
	return Vector2{
		x: v.x,
		y: newY,
	}
}

func (v Vector2) Floor() Vector2 {
	return Vector2{
		x: math.Floor(v.x),
		y: math.Floor(v.y),
	}
}

func (v Vector2) Dot(other Vector2) float64 {
	return (v.x * other.x) + (v.y * other.y)
}

// Perpendicular creates a vector perpendicular to the one passed in with the
// same magnitude
func (v Vector2) Perpendicular() Vector2 {
	return Vector2{
		x: v.y,
		y: -v.x,
	}
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

func (v Vector2) MultByVector(o Vector2) Vector2 {
	return Vector2{
		x: v.x * o.x,
		y: v.y * o.y,
	}
}

func (v Vector2) DivByConstant(t float64) Vector2 {
	return v.MultByConstant(1.0 / t)
}

// Distance is the euclidian distance between two points
func (v Vector2) Distance(other Vector2) float64 {
	return math.Sqrt(math.Pow(other.x-v.x, 2.0) + math.Pow(other.y-v.y, 2.0))
}
