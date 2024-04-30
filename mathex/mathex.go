package mathex

import (
	"math"

	"github.com/EliCDavis/vector"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float | vector.Number
}

func NearZero[T Number](v T) bool {
	const s = 1e-8
	return math.Abs(float64(v)) < s
}

func Npot[T Number](v T) T {
	r := 1
	for T(r) < v {
		r <<= 1
	}
	return T(r)
}

// FloatEquals - Check whether two given floats are almost equal
//func FloatEquals[FT FloatT](x, y FT) bool {
//	return (mathex.Abs(x-y) <= 0.000001*math.Max(1.0, math.Max(mathex.Abs(x), mathex.Abs(y))))
//}

// Lerp - Calculate linear interpolation between two floats
func Lerp[T Number](time float64, start, end T) T {
	return T(float64(start) + time*float64(end-start))
}

// Normalize - Normalize input value within input range
func Normalize[T Number](value, start, end T) T {
	return (value-start) / (end-start)
}

// Remap - Remap input value within input range to output range
func Remap[T Number](value, inputStart, inputEnd, outputStart, outputEnd T) T {
	return T(float64(value-inputStart)/float64(inputEnd-inputStart)*float64(outputEnd-outputStart) + float64(outputStart))
}

// Wrap - Wrap input value from min to max
func Wrap[T Number](value, min, max T) T {
	return T(float64(value) - float64(max-min)*Floor(float64(value-min)/float64(max-min)))
}

func Clamp[T Number](f, vmin, vmax T) T {
	return max(min(f, vmax), vmin)
}

func Abs[T Number](v T) T {
	return T(math.Abs(float64(v)))
}

func Round[T Number](v T) T {
	return T(math.Round(float64(v)))
}

func Ceil[T Number](v T) T {
	return T(math.Ceil(float64(v)))
}

func Floor[T Number](v T) T {
	return T(math.Floor(float64(v)))
}

func Sqrt[T Number](v T) T {
	return T(math.Sqrt(float64(v)))
}

func Cos[T constraints.Float](v T) T {
	return T(math.Cos(float64(v)))
}

func Sin[T constraints.Float](v T) T {
	return T(math.Sin(float64(v)))
}

func Acos[T constraints.Float](v T) T {
	return T(math.Acos(float64(v)))
}

func Asin[T constraints.Float](v T) T {
	return T(math.Asin(float64(v)))
}
