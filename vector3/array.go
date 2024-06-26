package vector3

import (
	"errors"
	"math"

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

func (v3a Array[T]) Add(other Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = Vector[T]{
			v.x + other.x,
			v.y + other.y,
			v.z + other.z,
		}
	}

	return
}

func (v3a Array[T]) AddInplace(other Vector[T]) Array[T] {
	for i, v := range v3a {
		v3a[i] = Vector[T]{
			v.x + other.x,
			v.y + other.y,
			v.z + other.z,
		}
	}
	return v3a
}

func (v3a Array[T]) Sub(other Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = Vector[T]{
			v.x - other.x,
			v.y - other.y,
			v.z - other.z,
		}
	}

	return
}

func (v3a Array[T]) SubInplace(other Vector[T]) Array[T] {
	for i, v := range v3a {
		v3a[i] = Vector[T]{
			v.x - other.x,
			v.y - other.y,
			v.z - other.z,
		}
	}
	return v3a
}

func (v3a Array[T]) Distance() (total float64) {
	if len(v3a) < 2 {
		return
	}
	for i := 1; i < len(v3a); i++ {
		total += v3a[i].Distance(v3a[i-1])
	}
	return
}

func (v3a Array[T]) Scale(t float64) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = Vector[T]{
			x: T(float64(v.x) * t),
			y: T(float64(v.y) * t),
			z: T(float64(v.z) * t),
		}
	}

	return
}

func (v3a Array[T]) ScaleInplace(t float64) Array[T] {
	for i, v := range v3a {
		v3a[i] = Vector[T]{
			x: T(float64(v.x) * t),
			y: T(float64(v.y) * t),
			z: T(float64(v.z) * t),
		}
	}
	return v3a
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

func (v3a Array[T]) ContainsNaN() bool {
	for _, v := range v3a {
		if v.ContainsNaN() {
			return true
		}
	}
	return false
}

func (v3a Array[T]) MaxLength() float64 {
	max := 0.

	for _, v := range v3a {
		max = math.Max(max, v.Length())
	}

	return max
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
func (v3a Array[T]) Average(vectors []Vector[T]) Vector[float64] {
	xTotal := 0.
	yTotal := 0.
	zTotal := 0.

	for _, v := range v3a {
		xTotal += float64(v.X())
		yTotal += float64(v.Y())
		zTotal += float64(v.Z())
	}

	return New(xTotal, yTotal, zTotal).DivByConstant(float64(len(v3a)))
}

// Bounds returns the min and max points of an AABB encompassing
func (v3a Array[T]) Bounds() (Vector[T], Vector[T]) {
	if len(v3a) == 0 {
		panic(errors.New("can not compute bounds from 0 vector elements"))
	}

	minV := v3a[0]
	maxV := v3a[0]

	for i := 1; i < len(v3a); i++ {
		v := v3a[i]
		minV.x = min(minV.x, v.x)
		minV.y = min(minV.y, v.y)
		minV.z = min(minV.z, v.z)

		maxV.x = max(maxV.x, v.x)
		maxV.y = max(maxV.y, v.y)
		maxV.z = max(maxV.z, v.z)
	}

	return minV, maxV
}

// StandardDeviation calculates the population standard deviation on each
// component of the vector
func (v3a Array[T]) StandardDeviation() (mean, deviation Vector[float64]) {
	mean = v3a.Average(v3a)

	xTotal, yTotal, zTotal := 0., 0., 0.
	for _, v := range v3a {
		diff := v.ToFloat64().Sub(mean)
		xTotal += (diff.x * diff.x)
		yTotal += (diff.y * diff.y)
		zTotal += (diff.z * diff.z)
	}

	deviation = New(
		math.Sqrt(xTotal/float64(len(v3a))),
		math.Sqrt(yTotal/float64(len(v3a))),
		math.Sqrt(zTotal/float64(len(v3a))),
	)
	return
}
