package mathex

import (
	"math"

	"github.com/EliCDavis/vector"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float | vector.Number
}

func Clamp[T Number](f, vmin, vmax T) T {
	return max(min(f, vmax), vmin)
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
