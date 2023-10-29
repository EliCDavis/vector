# Vector

![Coverage](https://img.shields.io/badge/Coverage-94.6%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/EliCDavis/vector)](https://goreportcard.com/report/github.com/EliCDavis/vector)
[![GoDoc](https://godoc.org/github.com/EliCDavis/vector?status.svg)](http://godoc.org/github.com/EliCDavis/vector)

Collection of **generic, immutable** vector math functions I've written overtime for different hobby projects.

Has support for both Vector2 (x, y), Vector3 (x, y, z), and Vector4 (x, y, z, w) functions.

## Example

Below is an example on how to implement the different sign distance field functions in a generic fashion to work for both int, int64, float32, and float64.

```go
package sdf

import (
	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector3"
)

type Field[T vector.Number] func(v vector3.Vector[T]) float64

func Sphere[T vector.Number](pos vector3.Vector[T], radius float64) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return v.Distance(pos) - radius
	}
}

func Box[T vector.Number](pos vector3.Vector[T], bounds vector3.Vector[T]) Field[T] {
	halfBounds := bounds.Scale(0.5)
	// It's best to watch the video to understand
	// https://www.youtube.com/watch?v=62-pRVZuS5c
	return func(v vector3.Vector[T]) float64 {
		q := v.Sub(pos).Abs().Sub(halfBounds)
		inside := math.Min(math.Max(q.X(), math.Max(q.Y(), q.Z())), 0)
		return vector3.Max(q, vector3.Zero[T]()).Length() + inside
	}
}

func Union[T vector.Number](a, b Field[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return math.Min(a(v), b(v))
	}
}

func Intersect[T vector.Number](a, b Field[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return math.Max(a(v), b(v))
	}
}

func Translate[T vector.Number](field Field[T], translation vector3.Vector[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return field(v.Sub(translation))
	}
}
```
