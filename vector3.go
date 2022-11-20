package vector

import (
	"math"
	"math/rand"
)

// Vector3 contains 3 components
type Vector3 struct {
	x float64
	y float64
	z float64
}

// NewVector3 creates a new vector with corresponding 3 components
func NewVector3(x float64, y float64, z float64) Vector3 {
	return Vector3{
		x: x,
		y: y,
		z: z,
	}
}

// Vector3Right is (1, 0, 0)
func Vector3Right() Vector3 {
	return NewVector3(1, 0, 0)
}

// Vector3Left is (-1, 0, 0)
func Vector3Left() Vector3 {
	return NewVector3(-1, 0, 0)
}

// Vector3Forward is (0, 0, 1)
func Vector3Forward() Vector3 {
	return NewVector3(0, 0, 1)
}

// Vector3Backwards is (0, 0, -1)
func Vector3Backwards() Vector3 {
	return NewVector3(0, 0, -1)
}

// Vector3Up is (0, 1, 0)
func Vector3Up() Vector3 {
	return NewVector3(0, 1, 0)
}

// Vector3Down is (0, -1, 0)
func Vector3Down() Vector3 {
	return NewVector3(0, -1, 0)
}

// Vector3Zero is (0, 0, 0)
func Vector3Zero() Vector3 {
	return NewVector3(0, 0, 0)
}

// Vector3One is (1, 1, 1)
func Vector3One() Vector3 {
	return NewVector3(1, 1, 1)
}

// AverageVector3 sums all vector3's components together and divides each
// component by the number of vectors added
func AverageVector3(vectors []Vector3) Vector3 {
	var center Vector3
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// X returns the x component
func (v Vector3) X() float64 {
	return v.x
}

// SetX changes the x component of the vector
func (v Vector3) SetX(newX float64) Vector3 {
	return Vector3{
		x: newX,
		y: v.y,
		z: v.z,
	}
}

// Y returns the y component
func (v Vector3) Y() float64 {
	return v.y
}

// SetY changes the y component of the vector
func (v Vector3) SetY(newY float64) Vector3 {
	return Vector3{
		x: v.x,
		y: newY,
		z: v.z,
	}
}

// Z returns the z component
func (v Vector3) Z() float64 {
	return v.z
}

// SetZ changes the z component of the vector
func (v Vector3) SetZ(newZ float64) Vector3 {
	return Vector3{
		x: v.x,
		y: v.y,
		z: newZ,
	}
}

// XY returns vector2 with the x and y components
func (v Vector3) XY() Vector2 {
	return NewVector2(v.x, v.y)
}

// XZ returns vector2 with the x and z components
func (v Vector3) XZ() Vector2 {
	return NewVector2(v.x, v.z)
}

// YZ returns vector2 with the y and z components
func (v Vector3) YZ() Vector2 {
	return NewVector2(v.y, v.z)
}

// Perpendicular finds a vector that meets this vector at a right angle.
// https://stackoverflow.com/a/11132720/4974261
func (v Vector3) Perpendicular() Vector3 {
	var c Vector3
	if v.Y() != 0 || v.Z() != 0 {
		c = Vector3Right()
	} else {
		c = Vector3Up()
	}
	return v.Cross(c)
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector3) Round() Vector3 {
	return NewVector3(
		math.Round(v.x),
		math.Round(v.y),
		math.Round(v.z),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Vector3) Floor() Vector3 {
	return NewVector3(
		math.Floor(v.x),
		math.Floor(v.y),
		math.Floor(v.z),
	)
}

// Ceil applies the ciel math operation to each component of the vector
func (v Vector3) Ceil() Vector3 {
	return NewVector3(
		math.Ceil(v.x),
		math.Ceil(v.y),
		math.Ceil(v.z),
	)
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v Vector3) Dot(other Vector3) float64 {
	return (v.x * other.x) + (v.y * other.y) + (v.z * other.z)
}

func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		x: (v.y * other.z) - (v.z * other.y),
		y: (v.z * other.x) - (v.x * other.z),
		z: (v.x * other.y) - (v.y * other.x),
	}
}

func (v Vector3) Normalized() Vector3 {
	return v.DivByConstant(v.Length())
}

func Vector3Rnd() Vector3 {
	return Vector3{
		x: rand.Float64(),
		y: rand.Float64(),
		z: rand.Float64(),
	}
}

func (v Vector3) MultByConstant(t float64) Vector3 {
	return Vector3{
		x: v.x * t,
		y: v.y * t,
		z: v.z * t,
	}
}

func (v Vector3) MultByVector(o Vector3) Vector3 {
	return Vector3{
		x: v.x * o.x,
		y: v.y * o.y,
		z: v.z * o.z,
	}
}

func (v Vector3) DivByConstant(t float64) Vector3 {
	return v.MultByConstant(1.0 / t)
}

func (v Vector3) Length() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y) + (v.z * v.z))
}

func (v Vector3) SquaredLength() float64 {
	return (v.x * v.x) + (v.y * v.y) + (v.z * v.z)
}

func (v Vector3) Distance(other Vector3) float64 {
	return math.Sqrt(math.Pow(other.x-v.x, 2.0) + math.Pow(other.y-v.y, 2.0) + math.Pow(other.z-v.z, 2.0))
}

func (v Vector3) Angle(other Vector3) float64 {
	denominator := math.Sqrt(v.SquaredLength() * other.SquaredLength())
	if denominator < 1e-15 {
		return 0.
	}
	return math.Acos(clamp(v.Dot(other)/denominator, -1., 1.))
}
