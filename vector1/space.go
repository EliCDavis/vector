package vector1

import (
	"math"

	"github.com/EliCDavis/vector"
)

type Space[T vector.Number] struct{}

func (Space[T]) Distance(a, b T) float64 {
	return math.Abs(float64(b - a))
}

func (Space[T]) Add(a, b T) T {
	return a + b
}

func (Space[T]) Sub(a, b T) T {
	return a - b
}

func (Space[T]) Scale(a T, amount float64) T {
	return T(float64(a) * amount)
}

func (Space[T]) Dot(a, b T) float64 {
	return float64(a * b)
}

func (Space[T]) Length(a T) float64 {
	return math.Abs(float64(a))
}

func (Space[T]) Normalized(a T) T {
	if a < 0 {
		return -1
	}
	if a == 0 {
		return 0
	}
	return 1
}

func (Space[T]) Lerp(a, b T, time float64) T {
	return T(float64(a+(b-a)) * time)
}
