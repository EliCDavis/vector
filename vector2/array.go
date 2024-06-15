package vector2

import (
	"errors"

	"github.com/EliCDavis/vector"
)

type Array[T vector.Number] []Vector[T]

type (
	Float64Array = Array[float64]
	Float32Array = Array[float32]
	IntArray     = Array[int]
	Int64Array   = Array[int64]
	Int32Array   = Array[int32]
	Int16Array   = Array[int16]
	Int8Array    = Array[int8]
)

func (v2a Array[T]) Add(other Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v2a))

	for i, v := range v2a {
		out[i] = Vector[T]{
			v.x + other.x,
			v.y + other.y,
		}
	}

	return
}

func (v2a Array[T]) AddInplace(other Vector[T]) Array[T] {
	for i, v := range v2a {
		v2a[i] = Vector[T]{
			v.x + other.x,
			v.y + other.y,
		}
	}
	return v2a
}

func (v2a Array[T]) Scale(t float64) (out Array[T]) {
	out = make(Array[T], len(v2a))

	for i, v := range v2a {
		out[i] = Vector[T]{
			x: T(float64(v.x) * t),
			y: T(float64(v.y) * t),
		}
	}

	return
}

func (v2a Array[T]) ScaleInplace(t float64) Array[T] {
	for i, v := range v2a {
		v2a[i] = Vector[T]{
			x: T(float64(v.x) * t),
			y: T(float64(v.y) * t),
		}
	}
	return v2a
}

func (v2a Array[T]) Normalized() (out Array[T]) {
	out = make(Array[T], len(v2a))

	for i, v := range v2a {
		out[i] = v.Normalized()
	}

	return
}

func (v2a Array[T]) Sum() (sum Vector[T]) {
	for _, v := range v2a {
		sum = sum.Add(v)
	}
	return
}

// Bounds returns the min and max points of an AABB encompassing
func (v2a Array[T]) Bounds() (Vector[T], Vector[T]) {
	if len(v2a) == 0 {
		panic(errors.New("can not compute bounds from 0 vector elements"))
	}

	minV := v2a[0]
	maxV := v2a[0]

	for i := 1; i < len(v2a); i++ {
		v := v2a[i]
		minV.x = min(minV.x, v.x)
		minV.y = min(minV.y, v.y)

		maxV.x = max(maxV.x, v.x)
		maxV.y = max(maxV.y, v.y)
	}

	return minV, maxV
}
