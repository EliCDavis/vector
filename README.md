# Vector

![Coverage](https://img.shields.io/badge/Coverage-96.7%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/EliCDavis/vector)](https://goreportcard.com/report/github.com/EliCDavis/vector)
[![GoDoc](https://godoc.org/github.com/EliCDavis/vector?status.svg)](http://godoc.org/github.com/EliCDavis/vector)

Collection of **generic, immutable** vector math functions I've written overtime for different hobby projects.

## API

| Function      | Vector2 | Vector3 | Vector4 | Description                                            |
|---------------|---------|---------|---------|--------------------------------------------------------|
| Abs           | ✅      | ✅     | ✅      | Returns a vector with each component's absolute value  |
| Add           | ✅      | ✅     | ✅      | Component Wise Addition                                |
| Angle         | ✅      | ✅     |         | Returns the angle between two vectors                  |
| Ceil          | ✅      | ✅     | ✅      | Ceils each vectors component to the nearest integer    |
| Clamp         | ✅      | ✅     | ✅      | Clamps each component between two values               |
| ContainsNaN   | ✅      | ✅     | ✅      | Returns true if any component of the vector is NaN     |
| Cross         |         | ✅     |          | Returns the cross product between two vectors          |
| Dot           | ✅      | ✅     | ✅      | Returns the dot product between two vectors            |
| Flip          | ✅      | ✅     | ✅      | Scales the vector by -1                                |
| FlipX         | ✅      | ✅     | ✅      | Returns a vector with the X component multiplied by -1 |
| FlipY         | ✅      | ✅     | ✅      | Returns a vector with the Y component multiplied by -1 |
| FlipZ         |         | ✅     | ✅      | Returns a vector with the Z component multiplied by -1 |
| FlipW         |         |         | ✅      | Returns a vector with the W component multiplied by -1 |
| Floor         | ✅      | ✅     | ✅      | Floors each vectors component                          |
| Format        | ✅      | ✅     | ✅      | Build a string with vector data                        |
| Length        | ✅      | ✅     | ✅      | Returns the length of the vector                       |
| LengthSquared | ✅      | ✅     | ✅      | Returns the squared length of the vector               |
| Max           | ✅      | ✅     | ✅      | Returns a new vector where each component is the largest value between the two vectors |
| MaxX          | ✅      | ✅     | ✅      | Returns the largest X component between the two vectors |
| MaxY          | ✅      | ✅     | ✅      | Returns the largest Y component between the two vectors |
| MaxZ          |         | ✅     | ✅      | Returns the largest Z component between the two vectors |
| MaxW          |         |         | ✅      | Returns the largest W component between the two vectors |
| MaxComponent  | ✅      | ✅     | ✅      | Returns the vectors largest component                  |
| Midpoint      | ✅      | ✅     | ✅      | Finds the mid point between two vectors                |
| Min           | ✅      | ✅     | ✅      | Returns a new vector where each component is the smallest value between the two vectors |
| MinX          | ✅      | ✅     | ✅      | Returns the smallest X component between the two vectors |
| MinY          | ✅      | ✅     | ✅      | Returns the smallest Y component between the two vectors |
| MinZ          |         | ✅     | ✅      | Returns the smallest Z component between the two vectors |
| MinW          |         |         | ✅      | Returns the smallest W component between the two vectors |
| MinComponent  | ✅      | ✅     | ✅      | Returns the vectors smallest component                 |
| Normalized    | ✅      | ✅     | ✅      | Returns the normalized vector                          |
| NearZero      | ✅      | ✅     | ✅      | Returns true if all of the components are near 0       |
| Round         | ✅      | ✅     | ✅      | Rounds each vectors component to the nearest integer   |
| Scale         | ✅      | ✅     | ✅      | Scales the vector by some constant                     |
| Sqrt          | ✅      | ✅     | ✅      | Returns a vector with each component's square root     |
| Sub           | ✅      | ✅     | ✅      | Component Wise Subtraction                             |
| Values        | ✅      | ✅     | ✅      | Returns all components of the vector                   |
| X             | ✅      | ✅     | ✅      | Returns the x component of the vector                  |
| Y             | ✅      | ✅     | ✅      | Returns the y component of the vector                  |
| Z             |         | ✅     | ✅      | Returns the z component of the vector                  |
| W             |         |         | ✅      | Returns the w component of the vector                  |
| XY            |         | ✅     | ✅      | Equivalent to vector2.New[T](v.x, v.y)                 |
| YZ            |         | ✅     | ✅      | Equivalent to vector2.New[T](v.y, v.z)                 |
| XZ            |         | ✅     | ✅      | Equivalent to vector2.New[T](v.x, v.z)                 |
| YX            | ✅      | ✅     | ✅      | Equivalent to vector2.New[T](v.y, v.x)                 |
| ZX            |         | ✅     | ✅      | Equivalent to vector2.New[T](v.z, v.x)                 |
| ZY            |         | ✅     | ✅      | Equivalent to vector2.New[T](v.z, v.y)                 |
| Log           | ✅      | ✅     | ✅      | Returns the natural logarithm for each component       |
| Log2          | ✅      | ✅     | ✅      | Returns the binary logarithm for each component        |
| Log10         | ✅      | ✅     | ✅      | Returns the decimal logarithm for each component       |
| Exp           | ✅      | ✅     | ✅      | Returns e**x, the base-e exponential for each component |
| Exp2          | ✅      | ✅     | ✅      | Returns 2**x, the base-2 exponential for each component |
| Expm1         | ✅      | ✅     | ✅      | Returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero |
| Write         | ✅      | ✅     | ✅      | Write vector component data as binary to io.Writer     |


## Example

Below is an example on how to implement the different sign distance field functions in a generic fashion to work for both `int8`, `int16`, `int32` `int`, `int64`, `float32`, and `float64`.

The code below produces this output:

![out.gif](./out.gif)


```go
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector2"
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
		inside := math.Min(float64(q.MaxComponent()), 0)
		return vector3.Max(q, vector3.Zero[T]()).Length() + inside
	}
}

func Union[T vector.Number](fields ...Field[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		min := math.MaxFloat64

		for _, f := range fields {
			fv := f(v)
			if fv < min {
				min = fv
			}
		}

		return min
	}
}

func Intersect[T vector.Number](fields ...Field[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		max := -math.MaxFloat64

		for _, f := range fields {
			fv := f(v)
			if fv > max {
				max = fv
			}
		}

		return max
	}
}

func Subtract[T vector.Number](minuend, subtrahend Field[T]) Field[T] {
	return func(f vector3.Vector[T]) float64 {
		return math.Max(minuend(f), -subtrahend(f))
	}
}

func Translate[T vector.Number](field Field[T], translation vector3.Vector[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return field(v.Sub(translation))
	}
}

func evaluateAtDepth[T vector.Number](dimension int, field Field[T], depth T, i int) {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{dimension, dimension},
	})

	for x := 0; x < dimension; x++ {
		for y := 0; y < dimension; y++ {
			v := field(vector3.New[T](T(x), T(y), depth))
			byteVal := (v / float64(dimension/20)) * 255
			var c color.Color
			if v > 0 {
				c = color.RGBA{R: 0, G: 0, B: byte(byteVal), A: 255}
			} else {
				c = color.RGBA{R: 0, G: byte(-byteVal), B: 0, A: 255}
			}
			img.Set(x, y, c)
		}
	}

	f, err := os.Create(fmt.Sprintf("field_%04d.png", i))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func main() {
	dimension := 512
	quarterDim := float64(dimension) / 4.

	middleCoord := vector2.
		Fill(dimension).
		Scale(0.5).
		ToFloat64()

	middleCord3D := vector3.New(middleCoord.X(), middleCoord.Y(), 0)

	smallRing := Subtract(
		Sphere(middleCord3D, 100),
		Sphere(middleCord3D, 50),
	)

	field := Intersect(
		Subtract(
			Sphere(middleCord3D, 200),
			Sphere(middleCord3D, 100),
		),
		Union(
			Translate(smallRing, vector3.New(1., 1., 0.).Scale(quarterDim)),
			Translate(smallRing, vector3.New(1., -1., 0.).Scale(quarterDim)),
			Translate(smallRing, vector3.New(-1., 1., 0.).Scale(quarterDim)),
			Translate(smallRing, vector3.New(-1., -1., 0.).Scale(quarterDim)),
			Box(middleCord3D, middleCord3D),
		),
	)

	for i := 0; i < 100; i++ {
		evaluateAtDepth(dimension, field, float64(-dimension/2)+(float64(i*dimension)*0.01), i)
	}
}

```

