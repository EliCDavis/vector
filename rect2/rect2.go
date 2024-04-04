package rect2

import (
	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector2"
)

type Rectangle[T vector.Number] struct {
	XY vector2.Vector[T]
	WH vector2.Vector[T]
}

type (
	Float64 = Rectangle[float64]
	Float32 = Rectangle[float32]
	Int     = Rectangle[int]
	Int64   = Rectangle[int64]
	Int32   = Rectangle[int32]
	Int16   = Rectangle[int16]
	Int8    = Rectangle[int8]
)

func New[T vector.Number](xy vector2.Vector[T], wh vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		XY: xy,
		WH: wh,
	}
}

func Zero[T vector.Number]() Rectangle[T] {
	return Rectangle[T]{
		XY: vector2.Zero[T](),
		WH: vector2.Zero[T](),
	}
}

func One[T vector.Number]() Rectangle[T] {
	return Rectangle[T]{
		XY: vector2.Zero[T](),
		WH: vector2.One[T](),
	}
}

func (r Rectangle[T]) A() vector2.Vector[T] {
	return r.XY
}

func (r Rectangle[T]) B() vector2.Vector[T] {
	return r.XY.Add(r.WH)
}

func (r Rectangle[T]) Delta(xy vector2.Vector[T], wh vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		XY: r.XY.Add(xy),
		WH: r.WH.Add(wh),
	}
}

func (r Rectangle[T]) Scale(f float64) Rectangle[T] {
	return Rectangle[T]{
		XY: r.XY,
		WH: r.WH.Scale(f),
	}
}

func (r Rectangle[T]) ScaleF(f float32) Rectangle[T] {
	return Rectangle[T]{
		XY: r.XY,
		WH: r.WH.ScaleF(f),
	}
}

func (r Rectangle[T]) ScaleByVector(f vector2.Float64) Rectangle[T] {
	return Rectangle[T]{
		XY: r.XY,
		WH: r.WH.ScaleByVector(f),
	}
}

func (r Rectangle[T]) ScaleByVectorF(f vector2.Float32) Rectangle[T] {
	return Rectangle[T]{
		XY: r.XY,
		WH: r.WH.ScaleByVectorF(f),
	}
}

func (r Rectangle[T]) Contains(v vector2.Vector[T]) bool {
	return vector2.GreaterEq(v, r.A()) && vector2.LessEq(v, r.B())
}
