package vector3

import (
	"math"

	"github.com/EliCDavis/vector"
)

type Array[T vector.Number] []Vector[T]

func (v3a Array[T]) Add(other Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = v.Add(other)
	}

	return
}

func (v3a Array[T]) Distance() float64 {
	if len(v3a) < 2 {
		return 0
	}
	total := 0.
	for i := 1; i < len(v3a); i++ {
		total += v3a[i].Distance(v3a[i-1])
	}
	return total
}

func (v3a Array[T]) MultByConstant(t float64) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = v.MultByConstant(t)
	}

	return
}

func (v3a Array[T]) DivByConstant(t float64) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = v.DivByConstant(t)
	}

	return
}

func (v3a Array[T]) Normalized() (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = v.Normalized()
	}

	return
}

func (v3a Array[T]) Sum() (sum Vector[T]) {
	for _, v := range v3a {
		sum = sum.Add(v)
	}
	return
}

func (v3a Array[T]) Modify(f func(Vector[T]) Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = f(v)
	}

	return
}

// Average sums all vector3's components together and divides each
// component by the number of values added
func (v3a Array[T]) Average(vectors []Vector[T]) Vector[T] {
	return v3a.Sum().DivByConstant(float64(len(vectors)))
}

// Bounds returns the min and max points of an AABB encompassing
func (v3a Array[T]) Bounds() (Vector[T], Vector[T]) {
	min := New(math.Inf(1), math.Inf(1), math.Inf(1))
	max := New(math.Inf(-1), math.Inf(-1), math.Inf(-1))

	for _, v := range v3a {
		min = New(
			math.Min(float64(v.x), min.x),
			math.Min(float64(v.y), min.y),
			math.Min(float64(v.z), min.z),
		)

		max = New(
			math.Max(float64(v.x), max.x),
			math.Max(float64(v.y), max.y),
			math.Max(float64(v.z), max.z),
		)
	}

	return New(T(min.x), T(min.y), T(min.z)), New(T(max.x), T(max.y), T(max.z))
}
