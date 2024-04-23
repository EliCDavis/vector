package vector2

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"

	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/mathex"
)

type Vector[T vector.Number] struct {
	X T
	Y T
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

func New[T vector.Number](x T, y T) Vector[T] {
	return Vector[T]{
		X: x,
		Y: y,
	}
}

// Fill creates a vector where each component is equal to v
func Fill[T vector.Number](v T) Vector[T] {
	return Vector[T]{
		X: v,
		Y: v,
	}
}

func Zero[T vector.Number]() Vector[T] {
	return Vector[T]{
		X: 0,
		Y: 0,
	}
}

func Up[T vector.Number]() Vector[T] {
	return Vector[T]{
		X: 0,
		Y: 1,
	}
}

func Down[T vector.Number]() Vector[T] {
	return Vector[T]{
		X: 0,
		Y: -1,
	}
}

func Left[T vector.Number]() Vector[T] {
	return Vector[T]{
		X: -1,
		Y: 0,
	}
}

func Right[T vector.Number]() Vector[T] {
	return Vector[T]{
		X: 1,
		Y: 0,
	}
}

func One[T vector.Number]() Vector[T] {
	return Vector[T]{
		X: 1,
		Y: 1,
	}
}

// Lerp linearly interpolates between a and b by t
func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {
	return Vector[T]{
		X: T((float64(b.X-a.X) * t) + float64(a.X)),
		Y: T((float64(b.Y-a.Y) * t) + float64(a.Y)),
	}
}

func Min[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		min(a.X, b.X),
		min(a.Y, b.Y),
	)
}

func Max[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		max(a.X, b.X),
		max(a.Y, b.Y),
	)
}

func MaxX[T vector.Number](a, b Vector[T]) T {
	return max(a.X, b.X)
}

func MaxY[T vector.Number](a, b Vector[T]) T {
	return max(a.Y, b.Y)
}

func MinX[T vector.Number](a, b Vector[T]) T {
	return min(a.X, b.X)
}

func MinY[T vector.Number](a, b Vector[T]) T {
	return min(a.Y, b.Y)
}

func Less[T vector.Number](a, b Vector[T]) bool {
	return a.X < b.X && a.Y < b.Y
}

func LessEq[T vector.Number](a, b Vector[T]) bool {
	return a.X <= b.X && a.Y <= b.Y
}

func Greater[T vector.Number](a, b Vector[T]) bool {
	return a.X > b.X && a.Y > b.Y
}

func GreaterEq[T vector.Number](a, b Vector[T]) bool {
	return a.X >= b.X && a.Y >= b.Y
}

func Midpoint[T vector.Number](a, b Vector[T]) Vector[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Vector[T]{
		X: T(float64(a.X+b.X) * 0.5),
		Y: T(float64(a.Y+b.Y) * 0.5),
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
		v.X = data[0]
	}

	if len(data) > 1 {
		v.Y = data[1]
	}

	return v
}

func Rand(r *rand.Rand) Vector[float64] {
	return Vector[float64]{
		X: r.Float64(),
		Y: r.Float64(),
	}
}

func (v Vector[T]) MinComponent() T {
	return min(v.X, v.Y)
}

func (v Vector[T]) MaxComponent() T {
	return max(v.X, v.Y)
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: float64(v.X),
		Y: float64(v.Y),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: 0,
		Y: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = T(aux.X)
	v.Y = T(aux.Y)
	return nil
}

func (v Vector[T]) Format(format string) string {
	return fmt.Sprintf(format, v.X, v.Y)
}

// Sqrt applies the math.Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		mathex.Sqrt(v.X),
		mathex.Sqrt(v.Y),
	)
}

func (v Vector[T]) Clamp(vmin, vmax T) Vector[T] {
	return Vector[T]{
		X: mathex.Clamp(v.X, vmin, vmax),
		Y: mathex.Clamp(v.Y, vmin, vmax),
	}
}

func (v Vector[T]) ClampV(vmin, vmax Vector[T]) Vector[T] {
	return Vector[T]{
		X: mathex.Clamp(v.X, vmin.X, vmax.X),
		Y: mathex.Clamp(v.Y, vmin.Y, vmax.Y),
	}
}

func (v Vector[T]) Clamp0V(vmax Vector[T]) Vector[T] {
	return Vector[T]{
		X: mathex.Clamp(v.X, 0, vmax.X),
		Y: mathex.Clamp(v.Y, 0, vmax.Y),
	}
}

func (v Vector[T]) ToNpot() Vector[T] {
	return Vector[T]{
		X: mathex.Npot(v.X),
		Y: mathex.Npot(v.Y),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return Vector[float64]{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return Vector[float32]{
		X: float32(v.X),
		Y: float32(v.Y),
	}
}

func (v Vector[T]) ToInt() Vector[int] {
	return Vector[int]{
		X: int(v.X),
		Y: int(v.Y),
	}
}

func (v Vector[T]) ToInt32() Vector[int32] {
	return Vector[int32]{
		X: int32(v.X),
		Y: int32(v.Y),
	}
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return Vector[int64]{
		X: int64(v.X),
		Y: int64(v.Y),
	}
}

func (v Vector[T]) Dx(dX T) Vector[T] {
	return Vector[T]{
		X: v.X + dX,
		Y: v.Y,
	}
}

func (v Vector[T]) Dy(dY T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y + dY,
	}
}

func (v Vector[T]) YX() Vector[T] {
	return Vector[T]{
		X: v.Y,
		Y: v.X,
	}
}

func (v Vector[T]) Angle(other Vector[T]) float64 {
	denominator := mathex.Sqrt((float64)(v.LengthSquared()) * (float64)(other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return math.Acos(mathex.Clamp((float64)(v.Dot(other))/denominator, -1., 1.))
}

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Vector[T]) Midpoint(o Vector[T]) Vector[T] {
	return o.Add(v).Scale(0.5)
}

func (v Vector[T]) Dot(other Vector[T]) T {
	return v.X*other.X + v.Y*other.Y
}

// Perpendicular creates a vector perpendicular to the one passed in with the
// same magnitude
func (v Vector[T]) Perpendicular() Vector[T] {
	return Vector[T]{
		X: v.Y,
		Y: -v.X,
	}
}

// Add returns a vector that is the result of two vectors added together
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector[T]) AddXY(x, y T) Vector[T] {
	return Vector[T]{
		X: v.X + x,
		Y: v.Y + y,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector[T]) SubXY(x, y T) Vector[T] {
	return Vector[T]{
		X: v.X - x,
		Y: v.Y - y,
	}
}

func (v Vector[T]) ReciprocalF() Vector[float32] {
	return Vector[float32]{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
	}
}

func (v Vector[T]) Reciprocal() Vector[float64] {
	return Vector[float64]{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
	}
}

func (v Vector[T]) Product() T {
	return v.X * v.Y
}

func (v Vector[T]) LengthSquared() T {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector[T]) Length() float64 {
	return math.Sqrt((float64)(v.LengthSquared()))
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) * t),
		Y: T(float64(v.Y) * t),
	}
}

func (v Vector[T]) ScaleF(t float32) Vector[T] {
	return Vector[T]{
		X: T(float32(v.X) * t),
		Y: T(float32(v.Y) * t),
	}
}

func (v Vector[T]) ScaleByVector(o Float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) * o.X),
		Y: T(float64(v.Y) * o.Y),
	}
}

func (v Vector[T]) ScaleByVectorF(o Float32) Vector[T] {
	return Vector[T]{
		X: T(float32(v.X) * o.X),
		Y: T(float32(v.Y) * o.Y),
	}
}

func (v Vector[T]) ScaleByVectorI(o Int) Vector[T] {
	return Vector[T]{
		X: v.X * T(o.X),
		Y: v.Y * T(o.Y),
	}
}

func (v Vector[T]) ScaleByXYF(x, y float32) Vector[T] {
	return Vector[T]{
		X: T(float32(v.X) * x),
		Y: T(float32(v.Y) * y),
	}
}

func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v Vector[T]) DivByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return v.Scale(1.0 / t)
}

func (v Vector[T]) DistanceSquared(other Vector[T]) T {
	xDist := other.X - v.X
	yDist := other.Y - v.Y
	return (xDist * xDist) + (yDist * yDist)
}

// Distance is the euclidean distance between two points
func (v Vector[T]) Distance(other Vector[T]) float64 {
	return math.Sqrt((float64)(v.DistanceSquared(other)))
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return Vector[T]{
		X: mathex.Round(v.X),
		Y: mathex.Round(v.Y),
	}
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(mathex.Round(v.X)),
		int(mathex.Round(v.Y)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return Vector[T]{
		X: mathex.Ceil(v.X),
		Y: mathex.Ceil(v.Y),
	}
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(mathex.Ceil(v.X)),
		int(mathex.Ceil(v.Y)),
	)
}

func (v Vector[T]) Floor() Vector[T] {
	return Vector[T]{
		X: mathex.Floor(v.X),
		Y: mathex.Floor(v.Y),
	}
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(mathex.Floor(v.X)),
		int(mathex.Floor(v.Y)),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return Vector[T]{
		X: T(mathex.Abs(v.X)),
		Y: T(mathex.Abs(v.Y)),
	}
}

func (v Vector[T]) NearZero() bool {
	return mathex.NearZero(v.X) && mathex.NearZero(v.Y)
}

func (v Vector[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.X)) {
		return true
	}

	if math.IsNaN(float64(v.Y)) {
		return true
	}

	return false
}

func (v Vector[T]) Flip() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y * -1,
	}
}

func (v Vector[T]) FlipX() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y,
	}
}

func (v Vector[T]) FlipY() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y * -1,
	}
}

func (v Vector[T]) Pivot(anchor Vector[T], wh Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X - wh.X*anchor.X,
		Y: v.Y - wh.Y*anchor.Y,
	}
}

// Log returns the natural logarithm for each component
func (v Vector[T]) Log() Vector[T] {
	return Vector[T]{
		X: T(math.Log(float64(v.X))),
		Y: T(math.Log(float64(v.Y))),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Vector[T]) Log10() Vector[T] {
	return Vector[T]{
		X: T(math.Log10(float64(v.X))),
		Y: T(math.Log10(float64(v.Y))),
	}
}

// Log2 returns the binary logarithm for each component
func (v Vector[T]) Log2() Vector[T] {
	return Vector[T]{
		X: T(math.Log2(float64(v.X))),
		Y: T(math.Log2(float64(v.Y))),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Vector[T]) Exp2() Vector[T] {
	return Vector[T]{
		X: T(math.Exp2(float64(v.X))),
		Y: T(math.Exp2(float64(v.Y))),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Vector[T]) Exp() Vector[T] {
	return Vector[T]{
		X: T(math.Exp(float64(v.X))),
		Y: T(math.Exp(float64(v.Y))),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Vector[T]) Expm1() Vector[T] {
	return Vector[T]{
		X: T(math.Expm1(float64(v.X))),
		Y: T(math.Expm1(float64(v.Y))),
	}
}

func (v Vector[T]) Values() (T, T) {
	return v.X, v.Y
}
