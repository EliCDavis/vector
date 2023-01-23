package vector2

import (
	"math"
	"math/rand"

	"github.com/EliCDavis/vector"
)

type Vector[T vector.Number] struct {
	x T
	y T
}

func New[T vector.Number](x T, y T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
	}
}

func Zero[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: 0,
	}
}

func Up[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: 1,
	}
}

func Down[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: -1,
	}
}

func Left[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: -1,
		y: 0,
	}
}

func Right[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 1,
		y: 0,
	}
}

func One[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 1,
		y: 1,
	}
}

func Rand() Vector[float64] {
	return Vector[float64]{
		x: rand.Float64(),
		y: rand.Float64(),
	}
}

func (v Vector[T]) X() T {
	return v.x
}

// SetX changes the x component of the vector
func (v Vector[T]) SetX(newX T) Vector[T] {
	return Vector[T]{
		x: newX,
		y: v.y,
	}
}

func (v Vector[T]) Y() T {
	return v.y
}

// SetY changes the y component of the vector
func (v Vector[T]) SetY(newY T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: newY,
	}
}

func (v Vector[T]) Floor() Vector[T] {
	return Vector[T]{
		x: T(math.Floor(float64(v.x))),
		y: T(math.Floor(float64(v.y))),
	}
}

func (v Vector[T]) Dot(other Vector[T]) float64 {
	return float64(v.x*other.x) + float64(v.y*other.y)
}

// Perpendicular creates a vector perpendicular to the one passed in with the
// same magnitude
func (v Vector[T]) Perpendicular() Vector[T] {
	return Vector[T]{
		x: v.y,
		y: -v.x,
	}
}

// Add returns a vector that is the result of two vectors added together
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v Vector[T]) Length() float64 {
	return math.Sqrt(float64(v.x*v.x) + float64(v.y*v.y))
}

func (v Vector[T]) Normalized() Vector[T] {
	if v.Length() == 0 {
		return New(v.x, v.y)
	}
	return v.DivByConstant(v.Length())
}

func (v Vector[T]) MultByConstant(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * t),
		y: T(float64(v.y) * t),
	}
}

func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x * o.x,
		y: v.y * o.y,
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return v.MultByConstant(1.0 / t)
}

// Distance is the euclidean distance between two points
func (v Vector[T]) Distance(other Vector[T]) float64 {
	return math.Sqrt(math.Pow(float64(other.x-v.x), 2.0) + math.Pow(float64(other.y-v.y), 2.0))
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return Vector[T]{
		x: T(math.Round(float64(v.x))),
		y: T(math.Round(float64(v.y))),
	}
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return Vector[T]{
		x: T(math.Ceil(float64(v.x))),
		y: T(math.Ceil(float64(v.y))),
	}
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return Vector[T]{
		x: T(math.Abs(float64(v.x))),
		y: T(math.Abs(float64(v.y))),
	}
}
