package vector3

import (
	"math"
	"math/rand"

	"github.com/EliCDavis/vector"
)

// Vector contains 3 components
type Vector[T vector.Number] struct {
	x T
	y T
	z T
}

type (
	Float64 = Vector[float64]
	Float32 = Vector[float32]
	Int     = Vector[int]
	Int64   = Vector[int64]
)

// New creates a new vector with corresponding 3 components
func New[T vector.Number](x T, y T, z T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
		z: z,
	}
}

// Right is (1, 0, 0)
func Right[T vector.Number]() Vector[T] {
	return New[T](1, 0, 0)
}

// Left is (-1, 0, 0)
func Left[T vector.Number]() Vector[T] {
	return New[T](-1, 0, 0)
}

// Forward is (0, 0, 1)
func Forward[T vector.Number]() Vector[T] {
	return New[T](0, 0, 1)
}

// Backwards is (0, 0, -1)
func Backwards[T vector.Number]() Vector[T] {
	return New[T](0, 0, -1)
}

// Up is (0, 1, 0)
func Up[T vector.Number]() Vector[T] {
	return New[T](0, 1, 0)
}

// Down is (0, -1, 0)
func Down[T vector.Number]() Vector[T] {
	return New[T](0, -1, 0)
}

// Zero is (0, 0, 0)
func Zero[T vector.Number]() Vector[T] {
	return New[T](0, 0, 0)
}

// One is (1, 1, 1)
func One[T vector.Number]() Vector[T] {
	return New[T](1, 1, 1)
}

// Average sums all vector3's components together and divides each
// component by the number of vectors added
func Average[T vector.Number](vectors []Vector[T]) Vector[T] {
	var center Vector[T]
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// X returns the x component
func (v Vector[T]) X() T {
	return v.x
}

// SetX changes the x component of the vector
func (v Vector[T]) SetX(newX T) Vector[T] {
	return Vector[T]{
		x: newX,
		y: v.y,
		z: v.z,
	}
}

// Y returns the y component
func (v Vector[T]) Y() T {
	return v.y
}

// SetX changes the x component of the vector
func (v Vector[T]) SetY(newY T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: newY,
		z: v.z,
	}
}

// Z returns the z component
func (v Vector[T]) Z() T {
	return v.z
}

// SetZ changes the x component of the vector
func (v Vector[T]) SetZ(newZ T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: newZ,
	}
}

// // XY returns vector2 with the x and y components
// func (v Vector3[T]) XY() Vector2 {
// 	return NewVector2(v.x, v.y)
// }

// // XZ returns vector2 with the x and z components
// func (v Vector3[T]) XZ() Vector2 {
// 	return NewVector2(v.x, v.z)
// }

// // YZ returns vector2 with the y and z components
// func (v Vector3[T]) YZ() Vector2 {
// 	return NewVector2(v.y, v.z)
// }

// Perpendicular finds a vector that meets this vector at a right angle.
// https://stackoverflow.com/a/11132720/4974261
func (v Vector[T]) Perpendicular() Vector[T] {
	var c Vector[T]
	if v.Y() != 0 || v.Z() != 0 {
		c = Right[T]()
	} else {
		c = Up[T]()
	}
	return v.Cross(c)
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return New(
		T(math.Round(float64(v.x))),
		T(math.Round(float64(v.y))),
		T(math.Round(float64(v.z))),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Vector[T]) Floor() Vector[T] {
	return New(
		T(math.Floor(float64(v.x))),
		T(math.Floor(float64(v.y))),
		T(math.Floor(float64(v.z))),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return New(
		T(math.Ceil(float64(v.x))),
		T(math.Ceil(float64(v.y))),
		T(math.Ceil(float64(v.z))),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return New(
		T(math.Abs(float64(v.x))),
		T(math.Abs(float64(v.y))),
		T(math.Abs(float64(v.z))),
	)
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v Vector[T]) Dot(other Vector[T]) float64 {
	return float64(v.x*other.x) + float64(v.y*other.y) + float64(v.z*other.z)
}

func (v Vector[T]) Cross(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: (v.y * other.z) - (v.z * other.y),
		y: (v.z * other.x) - (v.x * other.z),
		z: (v.x * other.y) - (v.y * other.x),
	}
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

func Rand() Vector[float64] {
	return Vector[float64]{
		x: rand.Float64(),
		y: rand.Float64(),
		z: rand.Float64(),
	}
}

func (v Vector[T]) MultByConstant(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * t),
		y: T(float64(v.y) * t),
		z: T(float64(v.z) * t),
	}
}

func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x * o.x,
		y: v.y * o.y,
		z: v.z * o.z,
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return v.MultByConstant(1.0 / t)
}

func (v Vector[T]) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vector[T]) SquaredLength() float64 {
	return math.Pow(float64(v.x), 2.0) + math.Pow(float64(v.y), 2.0) + math.Pow(float64(v.z), 2.0)
}

func (v Vector[T]) SquaredDistance(other Vector[T]) float64 {
	return math.Pow(float64(other.x-v.x), 2.0) + math.Pow(float64(other.y-v.y), 2.0) + math.Pow(float64(other.z-v.z), 2.0)
}

func (v Vector[T]) Distance(other Vector[T]) float64 {
	return math.Sqrt(v.SquaredDistance(other))
}

func (v Vector[T]) Angle(other Vector[T]) float64 {
	denominator := math.Sqrt(v.SquaredLength() * other.SquaredLength())
	if denominator < 1e-15 {
		return 0.
	}
	return math.Acos(clamp(v.Dot(other)/denominator, -1., 1.))
}
