# Vector
![Coverage](https://img.shields.io/badge/Coverage-66.0%25-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/EliCDavis/vector)](https://goreportcard.com/report/github.com/EliCDavis/vector)
[![GoDoc](https://godoc.org/github.com/EliCDavis/vector?status.svg)](http://godoc.org/github.com/EliCDavis/vector)

Collection of **generic, immutable** vector math functions I've written overtime for different hobby projects.

Has support for both Vector2 (x, y) and Vector3 (x, y, z) functions.

## Example

Below is an example on how to implement the sign distance field of a sphere in a generic fashion to work for both int, int64, float32, and float64.

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
```
