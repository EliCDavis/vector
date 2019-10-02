package vector

import "math"

type Vector3 struct {
	x float64
	y float64
	z float64
}

func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{
		x: x,
		y: y,
		z: z,
	}
}

func Vector3Zero() *Vector3 {
	return NewVector3(0, 0, 0)
}

func Vector3One() *Vector3 {
	return NewVector3(1, 1, 1)
}

func (v Vector3) X() float64 {
	return v.x
}

func (v Vector3) Y() float64 {
	return v.y
}

func (v Vector3) Z() float64 {
	return v.z
}

func (v Vector3) Add(other *Vector3) *Vector3 {
	return &Vector3{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v Vector3) Sub(other *Vector3) *Vector3 {
	return &Vector3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v Vector3) Dot(other *Vector3) float64 {
	return (v.x * other.x) + (v.y * other.y) + (v.z * other.z)
}

func (v Vector3) Cross(other *Vector3) *Vector3 {
	return &Vector3{
		x: (v.y * other.z) - (v.z * other.y),
		y: -((v.x * other.z) - (v.z * other.x)),
		z: (v.x * other.y) - (v.y * other.x),
	}
}

func (v Vector3) Normalized() *Vector3 {
	return v.DivByConstant(v.Length())
}

func (v Vector3) MultByConstant(t float64) *Vector3 {
	return &Vector3{
		x: v.x * t,
		y: v.y * t,
		z: v.z * t,
	}
}

func (v Vector3) MultByVector(o *Vector3) *Vector3 {
	return &Vector3{
		x: v.x * o.x,
		y: v.y * o.y,
		z: v.z * o.z,
	}
}

func (v Vector3) DivByConstant(t float64) *Vector3 {
	return v.MultByConstant(1.0 / t)
}

func (v Vector3) Length() float64 {
	return float64(math.Sqrt(float64((v.x * v.x) + (v.y * v.y) + (v.z * v.z))))
}

func (v Vector3) SquaredLength() float64 {
	return (v.x * v.x) + (v.y * v.y) + (v.z * v.z)
}

func (v Vector3) Distance(other *Vector3) float64 {
	return math.Sqrt(math.Pow(other.x-v.x, 2.0) + math.Pow(other.y-v.y, 2.0) + math.Pow(other.z-v.z, 2.0))
}
