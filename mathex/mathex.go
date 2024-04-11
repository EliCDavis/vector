package mathex

import (
	"math"

	"github.com/EliCDavis/vector"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float | vector.Number
}

// Lerp - Calculate linear interpolation between two floats
func Lerp[T Number](start, end, amount T) T {
	return start + amount*(end-start)
}

// Normalize - Normalize input value within input range
func Normalize[T Number](value, start, end T) T {
	return (value - start) / (end - start)
}

// Remap - Remap input value within input range to output range
func Remap[T Number](value, inputStart, inputEnd, outputStart, outputEnd T) T {
	return (value-inputStart)/(inputEnd-inputStart)*(outputEnd-outputStart) + outputStart
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
