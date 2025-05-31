package vector3

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector2"
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
	Int32   = Vector[int32]
	Int16   = Vector[int16]
	Int8    = Vector[int8]
)

// New creates a new vector with corresponding 3 components
func New[T vector.Number](x T, y T, z T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
		z: z,
	}
}

// Fill creates a vector where each component is equal to v
func Fill[T vector.Number](v T) Vector[T] {
	return Vector[T]{
		x: v,
		y: v,
		z: v,
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

func FromColor(c color.Color) Float64 {
	r, g, b, _ := c.RGBA()
	return New(float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff)
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

// Lerp linearly interpolates between a and b by t
func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {
	return Vector[T]{
		x: T((float64(b.x-a.x) * t) + float64(a.x)),
		y: T((float64(b.y-a.y) * t) + float64(a.y)),
		z: T((float64(b.z-a.z) * t) + float64(a.z)),
	}
}

func LerpClamped[T vector.Number](a, b Vector[T], t float64) Vector[T] {
	tClean := vector.Clamp(t, 0, 1)
	return Vector[T]{
		x: T((float64(b.x-a.x) * tClean) + float64(a.x)),
		y: T((float64(b.y-a.y) * tClean) + float64(a.y)),
		z: T((float64(b.z-a.z) * tClean) + float64(a.z)),
	}
}

func Min[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Min(float64(a.x), float64(b.x))),
		T(math.Min(float64(a.y), float64(b.y))),
		T(math.Min(float64(a.z), float64(b.z))),
	)
}

func Max[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Max(float64(a.x), float64(b.x))),
		T(math.Max(float64(a.y), float64(b.y))),
		T(math.Max(float64(a.z), float64(b.z))),
	)
}

func MaxX[T vector.Number](a, b Vector[T]) T {
	return T(math.Max(float64(a.x), float64(b.x)))
}

func MaxY[T vector.Number](a, b Vector[T]) T {
	return T(math.Max(float64(a.y), float64(b.y)))
}

func MaxZ[T vector.Number](a, b Vector[T]) T {
	return T(math.Max(float64(a.z), float64(b.z)))
}

func MinX[T vector.Number](a, b Vector[T]) T {
	return T(math.Min(float64(a.x), float64(b.x)))
}

func MinY[T vector.Number](a, b Vector[T]) T {
	return T(math.Min(float64(a.y), float64(b.y)))
}

func MinZ[T vector.Number](a, b Vector[T]) T {
	return T(math.Min(float64(a.z), float64(b.z)))
}

func Midpoint[T vector.Number](a, b Vector[T]) Vector[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Vector[T]{
		x: T(float64(a.x+b.x) * 0.5),
		y: T(float64(a.y+b.y) * 0.5),
		z: T(float64(a.z+b.z) * 0.5),
	}
}

// Builds a vector from the data found from the passed in array to the best of
// it's ability. If the length of the array is smaller than the vector itself,
// only those values will be used to build the vector, and the remaining vector
// components will remain the default value of the vector's data type (some
// version of 0).
func FromArray[T vector.Number](data []T) Vector[T] {
	v := Vector[T]{}

	if len(data) > 0 {
		v.x = data[0]
	}

	if len(data) > 1 {
		v.y = data[1]
	}

	if len(data) > 2 {
		v.z = data[2]
	}

	return v
}

func (v Vector[T]) ToArr() []T {
	return []T{v.x, v.y, v.z}
}

func (v Vector[T]) ToFixedArr() [3]T {
	return [3]T{v.x, v.y, v.z}
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{
		X: float64(v.x),
		Y: float64(v.y),
		Z: float64(v.z),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{
		X: 0,
		Y: 0,
		Z: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.x = T(aux.X)
	v.y = T(aux.Y)
	v.z = T(aux.Z)
	return nil
}

func (v Vector[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.x)) {
		return true
	}

	if math.IsNaN(float64(v.y)) {
		return true
	}

	if math.IsNaN(float64(v.z)) {
		return true
	}

	return false
}

func (v Vector[T]) Format(format string) string {
	return fmt.Sprintf(format, v.x, v.y, v.z)
}

func (v Vector[T]) MinComponent() T {
	return T(math.Min(float64(v.x), math.Min(float64(v.y), float64(v.z))))
}

func (v Vector[T]) MaxComponent() T {
	return T(math.Max(float64(v.x), math.Max(float64(v.y), float64(v.z))))
}

func (v Vector[T]) ToInt() Vector[int] {
	return Vector[int]{
		x: int(v.x),
		y: int(v.y),
		z: int(v.z),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return Vector[float64]{
		x: float64(v.x),
		y: float64(v.y),
		z: float64(v.z),
	}
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return Vector[float32]{
		x: float32(v.x),
		y: float32(v.y),
		z: float32(v.z),
	}
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return Vector[int64]{
		x: int64(v.x),
		y: int64(v.y),
		z: int64(v.z),
	}
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

// SetY changes the y component of the vector
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

// SetZ changes the z component of the vector
func (v Vector[T]) SetZ(newZ T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: newZ,
	}
}

func (v Vector[T]) XZY() Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.z,
		z: v.y,
	}
}

func (v Vector[T]) ZXY() Vector[T] {
	return Vector[T]{
		x: v.z,
		y: v.x,
		z: v.y,
	}
}

func (v Vector[T]) ZYX() Vector[T] {
	return Vector[T]{
		x: v.z,
		y: v.y,
		z: v.x,
	}
}

func (v Vector[T]) YXZ() Vector[T] {
	return Vector[T]{
		x: v.y,
		y: v.x,
		z: v.z,
	}
}

func (v Vector[T]) YZX() Vector[T] {
	return Vector[T]{
		x: v.y,
		y: v.z,
		z: v.x,
	}
}

// XY returns vector2 with the x and y components
func (v Vector[T]) XY() vector2.Vector[T] {
	return vector2.New(v.x, v.y)
}

// XZ returns vector2 with the x and z components
func (v Vector[T]) XZ() vector2.Vector[T] {
	return vector2.New(v.x, v.z)
}

// YZ returns vector2 with the y and z components
func (v Vector[T]) YZ() vector2.Vector[T] {
	return vector2.New(v.y, v.z)
}

// YX returns vector2 with the y and x components
func (v Vector[T]) YX() vector2.Vector[T] {
	return vector2.New(v.y, v.x)
}

// ZX returns vector2 with the z and x components
func (v Vector[T]) ZX() vector2.Vector[T] {
	return vector2.New(v.z, v.x)
}

// ZY returns vector2 with the z and y components
func (v Vector[T]) ZY() vector2.Vector[T] {
	return vector2.New(v.z, v.y)
}

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Vector[T]) Midpoint(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: T(float64(o.x+v.x) * 0.5),
		y: T(float64(o.y+v.y) * 0.5),
		z: T(float64(o.z+v.z) * 0.5),
	}
}

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

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(math.Round(float64(v.x))),
		int(math.Round(float64(v.y))),
		int(math.Round(float64(v.z))),
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

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(math.Floor(float64(v.x))),
		int(math.Floor(float64(v.y))),
		int(math.Floor(float64(v.z))),
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

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(math.Ceil(float64(v.x))),
		int(math.Ceil(float64(v.y))),
		int(math.Ceil(float64(v.z))),
	)
}

// Sqrt applies the math.Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		T(math.Sqrt(float64(v.x))),
		T(math.Sqrt(float64(v.y))),
		T(math.Sqrt(float64(v.z))),
	)
}

// Abs applies the Abs math operation to each component of the vector
//
//go:inline
func (v Vector[T]) Abs() Vector[T] {
	return New(
		T(math.Abs(float64(v.x))),
		T(math.Abs(float64(v.y))),
		T(math.Abs(float64(v.z))),
	)
}

//go:inline
func (v Vector[T]) Clamp(min, max T) Vector[T] {
	return Vector[T]{
		x: T(vector.Clamp(float64(v.x), float64(min), float64(max))),
		y: T(vector.Clamp(float64(v.y), float64(min), float64(max))),
		z: T(vector.Clamp(float64(v.z), float64(min), float64(max))),
	}
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
//
//go:inline
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
	return float64((v.x * other.x) + (v.y * other.y) + (v.z * other.z))
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

// Rand returns a vector with each component being a random value between [0.0, 1.0)
func Rand(r *rand.Rand) Vector[float64] {
	return Vector[float64]{
		x: r.Float64(),
		y: r.Float64(),
		z: r.Float64(),
	}
}

// RandRange returns a vector where each component is a random value that falls
// within the values of min and max
func RandRange[T vector.Number](r *rand.Rand, min, max T) Vector[T] {
	dist := float64(max - min)
	return Vector[T]{
		x: T(r.Float64()*dist) + min,
		y: T(r.Float64()*dist) + min,
		z: T(r.Float64()*dist) + min,
	}
}

// RandInUnitSphere returns a randomly sampled point in or on the unit
func RandInUnitSphere(r *rand.Rand) Vector[float64] {
	for {
		p := RandRange(r, -1., 1.)
		if p.LengthSquared() < 1 {
			return p
		}
	}
}

// RandNormal returns a random normal
func RandNormal(r *rand.Rand) Vector[float64] {
	return Vector[float64]{
		x: -1. + (r.Float64() * 2.),
		y: -1. + (r.Float64() * 2.),
		z: -1. + (r.Float64() * 2.),
	}.Normalized()
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * t),
		y: T(float64(v.y) * t),
		z: T(float64(v.z) * t),
	}
}

func (v Vector[T]) Reflect(normal Vector[T]) Vector[T] {
	return v.Sub(normal.Scale(2. * v.Dot(normal)))
}

func (v Vector[T]) Refract(normal Vector[T], etaiOverEtat float64) Vector[T] {
	cosTheta := math.Min(v.Scale(-1).Dot(normal), 1.0)
	perpendicular := v.Add(normal.Scale(cosTheta)).Scale(etaiOverEtat)
	parallel := normal.Scale(-math.Sqrt(math.Abs(1.0 - perpendicular.LengthSquared())))
	return perpendicular.Add(parallel)
}

// MultByVector is component wise multiplication, also known as Hadamard product.
func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x * o.x,
		y: v.y * o.y,
		z: v.z * o.z,
	}
}

//go:inline
func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) / t),
		y: T(float64(v.y) / t),
		z: T(float64(v.z) / t),
	}
}

func (v Vector[T]) DivByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x / o.x,
		y: v.y / o.y,
		z: v.z / o.z,
	}
}

func (v Vector[T]) Mod(t float64) Vector[T] {
	return Vector[T]{
		x: T(math.Mod(float64(v.x), t)),
		y: T(math.Mod(float64(v.y), t)),
		z: T(math.Mod(float64(v.z), t)),
	}
}

func (v Vector[T]) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vector[T]) LengthSquared() float64 {
	return float64((v.x * v.x) + (v.y * v.y) + (v.z * v.z))
}

func (v Vector[T]) DistanceSquared(other Vector[T]) float64 {
	xDist := other.x - v.x
	yDist := other.y - v.y
	zDist := other.z - v.z
	return float64((xDist * xDist) + (yDist * yDist) + (zDist * zDist))
}

func (v Vector[T]) Distance(other Vector[T]) float64 {
	return math.Sqrt(v.DistanceSquared(other))
}

func (v Vector[T]) Angle(other Vector[T]) float64 {
	denominator := math.Sqrt(v.LengthSquared() * other.LengthSquared())
	if denominator < 1e-15 {
		return 0.
	}
	return math.Acos(vector.Clamp(v.Dot(other)/denominator, -1., 1.))
}

func (v Vector[T]) NearZero() bool {
	const s = 1e-8
	return (math.Abs(float64(v.x)) < s) && (math.Abs(float64(v.y)) < s) && (math.Abs(float64(v.z)) < s)
}

func (v Vector[T]) Flip() Vector[T] {
	return Vector[T]{
		x: v.x * -1,
		y: v.y * -1,
		z: v.z * -1,
	}
}

func (v Vector[T]) FlipX() Vector[T] {
	return Vector[T]{
		x: v.x * -1,
		y: v.y,
		z: v.z,
	}
}

func (v Vector[T]) FlipY() Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y * -1,
		z: v.z,
	}
}

func (v Vector[T]) FlipZ() Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: v.z * -1,
	}
}

// Log returns the natural logarithm for each component
func (v Vector[T]) Log() Vector[T] {
	return Vector[T]{
		x: T(math.Log(float64(v.x))),
		y: T(math.Log(float64(v.y))),
		z: T(math.Log(float64(v.z))),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Vector[T]) Log10() Vector[T] {
	return Vector[T]{
		x: T(math.Log10(float64(v.x))),
		y: T(math.Log10(float64(v.y))),
		z: T(math.Log10(float64(v.z))),
	}
}

// Log2 returns the binary logarithm for each component
func (v Vector[T]) Log2() Vector[T] {
	return Vector[T]{
		x: T(math.Log2(float64(v.x))),
		y: T(math.Log2(float64(v.y))),
		z: T(math.Log2(float64(v.z))),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Vector[T]) Exp2() Vector[T] {
	return Vector[T]{
		x: T(math.Exp2(float64(v.x))),
		y: T(math.Exp2(float64(v.y))),
		z: T(math.Exp2(float64(v.z))),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Vector[T]) Exp() Vector[T] {
	return Vector[T]{
		x: T(math.Exp(float64(v.x))),
		y: T(math.Exp(float64(v.y))),
		z: T(math.Exp(float64(v.z))),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Vector[T]) Expm1() Vector[T] {
	return Vector[T]{
		x: T(math.Expm1(float64(v.x))),
		y: T(math.Expm1(float64(v.y))),
		z: T(math.Expm1(float64(v.z))),
	}
}

func (v Vector[T]) Values() (T, T, T) {
	return v.x, v.y, v.z
}

func (v Vector[T]) Reciprocal() Vector[float64] {
	return Vector[float64]{
		x: 1.0 / float64(v.x),
		y: 1.0 / float64(v.y),
		z: 1.0 / float64(v.z),
	}
}

func (v Vector[T]) Component(index int) T {
	switch index {
	case 0:
		return v.x

	case 1:
		return v.y

	case 2:
		return v.z

	default:
		panic(fmt.Errorf("invalid index: %d", index))
	}
}
