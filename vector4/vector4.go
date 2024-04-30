package vector4

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"

	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/mathex"
	"github.com/EliCDavis/vector/vector2"
	"github.com/EliCDavis/vector/vector3"
)

// Vector contains 4 components
type Vector[T vector.Number] struct {
	x T
	y T
	z T
	w T
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
func New[T vector.Number](x, y, z, w T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

// Fill creates a vector where each component is equal to v
func Fill[T vector.Number](v T) Vector[T] {
	return Vector[T]{
		x: v,
		y: v,
		z: v,
		w: v,
	}
}

func FromColor(c color.Color) Float64 {
	r, g, b, a := c.RGBA()
	return New(float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff, float64(a)/0xffff)
}

// Zero is (0, 0, 0)
func Zero[T vector.Number]() Vector[T] {
	return New[T](0, 0, 0, 0)
}

// One is (1, 1, 1)
func One[T vector.Number]() Vector[T] {
	return New[T](1, 1, 1, 1)
}

// Average sums all vector4's components together and divides each
// component by the number of vectors added
func Average[T vector.Number](vectors []Vector[T]) Vector[T] {
	var center Vector[T]
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// Lerp linearly interpolates between a and b by t
// func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {

// 	// (b - a) * t + a
// 	// bt - at + a
// 	// bt - a(1 - t)
// 	tm1 := 1. - t
// 	return Vector[T]{
// 		x: T((float64(b.x) * t) - (float64(a.x) * tm1)),
// 		y: T((float64(b.y) * t) - (float64(a.y) * tm1)),
// 		z: T((float64(b.z) * t) - (float64(a.z) * tm1)),
// 		w: T((float64(b.w) * t) - (float64(a.w) * tm1)),
// 	}
// }

// Lerp linearly interpolates between a and b by t
func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {

	// return b.Sub(a).Scale(t).Add(a)
	return Vector[T]{
		x: T((float64(b.x-a.x) * t) + float64(a.x)),
		y: T((float64(b.y-a.y) * t) + float64(a.y)),
		z: T((float64(b.z-a.z) * t) + float64(a.z)),
		w: T((float64(b.w-a.w) * t) + float64(a.w)),
	}
}

func (v Vector[T]) Negated() Vector[T] {
	return Vector[T]{
		x: -v.x,
		y: -v.y,
		z: -v.z,
		w: -v.w,
	}
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * t),
		y: T(float64(v.y) * t),
		z: T(float64(v.z) * t),
		w: T(float64(v.w) * t),
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) / t),
		y: T(float64(v.y) / t),
		z: T(float64(v.z) / t),
		w: T(float64(v.w) / t),
	}
}

func Min[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		min(a.x, b.x),
		min(a.y, b.y),
		min(a.z, b.z),
		min(a.w, b.w),
	)
}

func Max[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		max(a.x, b.x),
		max(a.y, b.y),
		max(a.z, b.z),
		max(a.w, b.w),
	)
}

func MaxX[T vector.Number](a, b Vector[T]) T {
	return max(a.x, b.x)
}

func MaxY[T vector.Number](a, b Vector[T]) T {
	return max(a.y, b.y)
}

func MaxZ[T vector.Number](a, b Vector[T]) T {
	return max(a.z, b.z)
}

func MaxW[T vector.Number](a, b Vector[T]) T {
	return max(a.w, b.w)
}

func MinX[T vector.Number](a, b Vector[T]) T {
	return min(a.x, b.x)
}

func MinY[T vector.Number](a, b Vector[T]) T {
	return min(a.y, b.y)
}

func MinZ[T vector.Number](a, b Vector[T]) T {
	return min(a.z, b.z)
}

func MinW[T vector.Number](a, b Vector[T]) T {
	return min(a.w, b.w)
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
		w: T(float64(a.w+b.w) * 0.5),
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

	if len(data) > 3 {
		v.w = data[3]
	}

	return v
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: float64(v.x),
		Y: float64(v.y),
		Z: float64(v.z),
		W: float64(v.w),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.x = T(aux.X)
	v.y = T(aux.Y)
	v.z = T(aux.Z)
	v.w = T(aux.W)
	return nil
}

func (v Vector[T]) Format(format string) string {
	return fmt.Sprintf(format, v.x, v.y, v.z, v.w)
}

func (v Vector[T]) MinComponent() T {
	return min(v.x, v.y, v.z, v.w)
}

func (v Vector[T]) MaxComponent() T {
	return max(v.x, v.y, v.z, v.w)
}

func (v Vector[T]) ToInt() Vector[int] {
	return Vector[int]{
		x: int(v.x),
		y: int(v.y),
		z: int(v.z),
		w: int(v.w),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return Vector[float64]{
		x: float64(v.x),
		y: float64(v.y),
		z: float64(v.z),
		w: float64(v.w),
	}
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return Vector[float32]{
		x: float32(v.x),
		y: float32(v.y),
		z: float32(v.z),
		w: float32(v.w),
	}
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return Vector[int64]{
		x: int64(v.x),
		y: int64(v.y),
		z: int64(v.z),
		w: int64(v.w),
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
		w: v.w,
	}
}

func (v Vector[T]) AddX(dX T) Vector[T] {
	return Vector[T]{
		x: v.x + dX,
		y: v.y,
		z: v.z,
		w: v.w,
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
		w: v.w,
	}
}

func (v Vector[T]) AddY(dY T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y + dY,
		z: v.z,
		w: v.w,
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
		w: v.w,
	}
}

func (v Vector[T]) AddZ(dZ T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: v.z + dZ,
		w: v.w,
	}
}

// W returns the w component
func (v Vector[T]) W() T {
	return v.w
}

// SetW changes the w component of the vector
func (v Vector[T]) SetW(newW T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: v.z,
		w: newW,
	}
}

func (v Vector[T]) AddW(dW T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: v.z,
		w: v.w + dW,
	}
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
		w: v.w + other.w,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
		w: v.w - other.w,
	}
}

func (v Vector[T]) ReciprocalF() Vector[float32] {
	return Vector[float32]{
		x: 1.0 / float32(v.x),
		y: 1.0 / float32(v.y),
		z: 1.0 / float32(v.z),
		w: 1.0 / float32(v.w),
	}
}

func (v Vector[T]) Reciprocal() Vector[float64] {
	return Vector[float64]{
		x: 1.0 / float64(v.x),
		y: 1.0 / float64(v.y),
		z: 1.0 / float64(v.z),
		w: 1.0 / float64(v.w),
	}
}

func (v Vector[T]) Product() T {
	return v.x * v.y * v.z * v.w
}

func (v Vector[T]) Dot(other Vector[T]) float64 {
	return float64((v.x * other.x) + (v.y * other.y) + (v.z * other.z) + (v.w * other.w))
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

func (v Vector[T]) Length() float64 {
	return mathex.Sqrt(float64(v.LengthSquared()))
}

func (v Vector[T]) LengthF() float32 {
	return mathex.Sqrt(float32(v.LengthSquared()))
}

func (v Vector[T]) LengthSquared() T {
	return (v.x * v.x) + (v.y * v.y) + (v.z * v.z) + (v.w * v.w)
}

// Sqrt applies the math.Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		mathex.Sqrt(v.x),
		mathex.Sqrt(v.y),
		mathex.Sqrt(v.z),
		mathex.Sqrt(v.w),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return New(
		mathex.Abs(v.x),
		mathex.Abs(v.y),
		mathex.Abs(v.z),
		mathex.Abs(v.w),
	)
}

func (v Vector[T]) Clamp(vmin, vmax T) Vector[T] {
	return Vector[T]{
		x: mathex.Clamp(v.x, vmin, vmax),
		y: mathex.Clamp(v.y, vmin, vmax),
		z: mathex.Clamp(v.z, vmin, vmax),
		w: mathex.Clamp(v.w, vmin, vmax),
	}
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return New(
		mathex.Round(v.x),
		mathex.Round(v.y),
		mathex.Round(v.z),
		mathex.Round(v.w),
	)
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(mathex.Round(v.x)),
		int(mathex.Round(v.y)),
		int(mathex.Round(v.z)),
		int(mathex.Round(v.w)),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Vector[T]) Floor() Vector[T] {
	return New(
		mathex.Floor(v.x),
		mathex.Floor(v.y),
		mathex.Floor(v.z),
		mathex.Floor(v.w),
	)
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(mathex.Floor(v.x)),
		int(mathex.Floor(v.y)),
		int(mathex.Floor(v.z)),
		int(mathex.Floor(v.w)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return New(
		mathex.Ceil(v.x),
		mathex.Ceil(v.y),
		mathex.Ceil(v.z),
		mathex.Ceil(v.w),
	)
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(mathex.Ceil(v.x)),
		int(mathex.Ceil(v.y)),
		int(mathex.Ceil(v.z)),
		int(mathex.Ceil(v.w)),
	)
}

// MultByVector is component wise multiplication, also known as Hadamard product.
func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x * o.x,
		y: v.y * o.y,
		z: v.z * o.z,
		w: v.w * o.w,
	}
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

	if math.IsNaN(float64(v.w)) {
		return true
	}

	return false
}

func (v Vector[T]) NearZero() bool {
	return mathex.NearZero(v.x) && mathex.NearZero(v.y) && mathex.NearZero(v.z) && mathex.NearZero(v.w)
}

func (v Vector[T]) Flip() Vector[T] {
	return Vector[T]{
		x: v.x * -1,
		y: v.y * -1,
		z: v.z * -1,
		w: v.w * -1,
	}
}

func (v Vector[T]) FlipX() Vector[T] {
	return Vector[T]{
		x: v.x * -1,
		y: v.y,
		z: v.z,
		w: v.w,
	}
}

func (v Vector[T]) FlipY() Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y * -1,
		z: v.z,
		w: v.w,
	}
}

func (v Vector[T]) FlipZ() Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: v.z * -1,
		w: v.w,
	}
}

func (v Vector[T]) FlipW() Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: v.z,
		w: v.w * -1,
	}
}

func (v Vector[T]) XYZ() vector3.Vector[T] {
	return vector3.New[T](v.x, v.y, v.z)
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

// Log returns the natural logarithm for each component
func (v Vector[T]) Log() Vector[T] {
	return Vector[T]{
		x: T(math.Log(float64(v.x))),
		y: T(math.Log(float64(v.y))),
		z: T(math.Log(float64(v.z))),
		w: T(math.Log(float64(v.w))),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Vector[T]) Log10() Vector[T] {
	return Vector[T]{
		x: T(math.Log10(float64(v.x))),
		y: T(math.Log10(float64(v.y))),
		z: T(math.Log10(float64(v.z))),
		w: T(math.Log10(float64(v.w))),
	}
}

// Log2 returns the binary logarithm for each component
func (v Vector[T]) Log2() Vector[T] {
	return Vector[T]{
		x: T(math.Log2(float64(v.x))),
		y: T(math.Log2(float64(v.y))),
		z: T(math.Log2(float64(v.z))),
		w: T(math.Log2(float64(v.w))),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Vector[T]) Exp2() Vector[T] {
	return Vector[T]{
		x: T(math.Exp2(float64(v.x))),
		y: T(math.Exp2(float64(v.y))),
		z: T(math.Exp2(float64(v.z))),
		w: T(math.Exp2(float64(v.w))),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Vector[T]) Exp() Vector[T] {
	return Vector[T]{
		x: T(math.Exp(float64(v.x))),
		y: T(math.Exp(float64(v.y))),
		z: T(math.Exp(float64(v.z))),
		w: T(math.Exp(float64(v.w))),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Vector[T]) Expm1() Vector[T] {
	return Vector[T]{
		x: T(math.Expm1(float64(v.x))),
		y: T(math.Expm1(float64(v.y))),
		z: T(math.Expm1(float64(v.z))),
		w: T(math.Expm1(float64(v.w))),
	}
}

func (v Vector[T]) Values() (T, T, T, T) {
	return v.x, v.y, v.z, v.w
}
