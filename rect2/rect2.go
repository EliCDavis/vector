package rect2

import (
	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector2"
)

type Rectangle[T vector.Number] struct {
	position vector2.Vector[T]
	size     vector2.Vector[T]
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

func New[T vector.Number](position vector2.Vector[T], size vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		position: position,
		size:     size,
	}
}

func Zero[T vector.Number]() Rectangle[T] {
	return Rectangle[T]{
		position: vector2.Zero[T](),
		size:     vector2.Zero[T](),
	}
}

func One[T vector.Number]() Rectangle[T] {
	return Rectangle[T]{
		position: vector2.Zero[T](),
		size:     vector2.One[T](),
	}
}

func (r Rectangle[T]) A() vector2.Vector[T] {
	return r.position
}

func (r Rectangle[T]) SetA(a vector2.Vector[T]) Rectangle[T] {
	dxy := a.Sub(r.position)
	return Rectangle[T]{
		position: a,
		size:     r.size.Sub(dxy),
	}
}

func (r Rectangle[T]) B() vector2.Vector[T] {
	return r.position.Add(r.size)
}

func (r Rectangle[T]) SetB(b vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     b,
	}
}

func (r Rectangle[T]) HorizontalLine(y T) (vector2.Vector[T], vector2.Vector[T]) {
	return vector2.New(r.A().X(), y), vector2.New(r.B().X(), y)
}

func (r Rectangle[T]) VerticalLine(x T) (vector2.Vector[T], vector2.Vector[T]) {
	return vector2.New(x, r.A().Y()), vector2.New(x, r.B().Y())
}

func (r Rectangle[T]) Center() vector2.Vector[T] {
	return r.position.Add(r.size.ScaleF(0.5))
}

func (v Rectangle[T]) ToFloat64() Rectangle[float64] {
	return Rectangle[float64]{
		position: v.position.ToFloat64(),
		size:     v.size.ToFloat64(),
	}
}

func (v Rectangle[T]) ToFloat32() Rectangle[float32] {
	return Rectangle[float32]{
		position: v.position.ToFloat32(),
		size:     v.size.ToFloat32(),
	}
}

func (v Rectangle[T]) ToInt() Rectangle[int] {
	return Rectangle[int]{
		position: v.position.ToInt(),
		size:     v.size.ToInt(),
	}
}

func (v Rectangle[T]) ToInt32() Rectangle[int32] {
	return Rectangle[int32]{
		position: v.position.ToInt32(),
		size:     v.size.ToInt32(),
	}
}

func (v Rectangle[T]) ToInt64() Rectangle[int64] {
	return Rectangle[int64]{
		position: v.position.ToInt64(),
		size:     v.size.ToInt64(),
	}
}

// X returns the x of the xy component
func (r Rectangle[T]) X() T {
	return r.position.X()
}

// SetX changes the x of the xy component of the rectangle
func (r Rectangle[T]) SetX(newX T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.SetX(newX),
		size:     r.size,
	}
}

func (r Rectangle[T]) AddX(dX T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.AddX(dX),
		size:     r.size,
	}
}

// Y returns the y of the xy component
func (r Rectangle[T]) Y() T {
	return r.position.Y()
}

// SetY changes the y of the xy component of the rectangle
func (r Rectangle[T]) SetY(newY T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.SetY(newY),
		size:     r.size,
	}
}

func (r Rectangle[T]) AddY(dY T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.AddY(dY),
		size:     r.size,
	}
}

// Width returns the x of the wh component
func (r Rectangle[T]) Width() T {
	return r.size.X()
}

// SetWidth changes the x of the wh component of the rectangle
func (r Rectangle[T]) SetWidth(newW T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.SetX(newW),
	}
}

func (r Rectangle[T]) AddWidth(dW T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.AddX(dW),
	}
}

// Y returns the y of the wh component
func (r Rectangle[T]) Height() T {
	return r.size.Y()
}

// SetHeight changes the y of the wh component of the rectangle
func (r Rectangle[T]) SetHeight(newH T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.SetY(newH),
	}
}

func (r Rectangle[T]) AddHeight(dH T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.AddY(dH),
	}
}

// XY returns the xy component
func (r Rectangle[T]) Position() vector2.Vector[T] {
	return r.position
}

// SetPosition changes the xy component of the rectangle
func (r Rectangle[T]) SetPosition(newXY vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		position: newXY,
		size:     r.size,
	}
}

// SetPosition changes the xy component of the rectangle
func (r Rectangle[T]) SetPositionXY(x, y T) Rectangle[T] {
	return Rectangle[T]{
		position: vector2.New(x, y),
		size:     r.size,
	}
}

func (r Rectangle[T]) AddPosition(dXY vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.Add(dXY),
		size:     r.size,
	}
}

func (r Rectangle[T]) AddPositionXY(x, y T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.AddXY(x, y),
		size:     r.size,
	}
}

// Size returns the wh component
func (r Rectangle[T]) Size() vector2.Vector[T] {
	return r.size
}

// SetSize changes the wh component of the rectangle
func (r Rectangle[T]) SetSize(newWH vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     newWH,
	}
}

// SetSizeXY changes the wh component of the rectangle
func (r Rectangle[T]) SetSizeXY(width, height T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     vector2.New(width, height),
	}
}

func (r Rectangle[T]) AddSize(dWH vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.Add(dWH),
	}
}

func (r Rectangle[T]) AddSizeXY(width, height T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.AddXY(width, height),
	}
}

// Round takes each component of the rectangle and rounds it to the nearest whole
// number
func (v Rectangle[T]) Round() Rectangle[T] {
	return New(
		v.position.Round(),
		v.size.Round(),
	)
}

// RoundToInt takes each component of the rectangle and rounds it to the nearest
// whole number, and then casts it to a int
func (v Rectangle[T]) RoundToInt() Rectangle[int] {
	return New(
		v.position.RoundToInt(),
		v.size.RoundToInt(),
	)
}

// Ceil applies the ceil math operation to each component of the rectangle
func (v Rectangle[T]) Ceil() Rectangle[T] {
	return New(
		v.position.Ceil(),
		v.size.Ceil(),
	)
}

// CeilToInt applies the ceil math operation to each component of the rectangle,
// and then casts it to a int
func (v Rectangle[T]) CeilToInt() Rectangle[int] {
	return New(
		v.position.CeilToInt(),
		v.size.CeilToInt(),
	)
}

// Floor applies the floor math operation to each component of the rectangle
func (v Rectangle[T]) Floor() Rectangle[T] {
	return New(
		v.position.Floor(),
		v.size.Floor(),
	)
}

// FloorToInt applies the floor math operation to each component of the rectangle,
// and then casts it to a int
func (v Rectangle[T]) FloorToInt() Rectangle[int] {
	return New(
		v.position.FloorToInt(),
		v.size.FloorToInt(),
	)
}

func (r Rectangle[T]) Add(xy vector2.Vector[T], wh vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.Add(xy),
		size:     r.size.Add(wh),
	}
}

func (r Rectangle[T]) AddXYWH(x, y, w, h T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.AddXY(x, y),
		size:     r.size.AddXY(w, h),
	}
}

func (r Rectangle[T]) ShrinkXYWH(left, top, right, bottom T) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.AddXY(left, top),
		size:     r.size.AddXY(-left-right, -top-bottom),
	}
}

func (r Rectangle[T]) Scale(f float64) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.Scale(f),
	}
}

func (r Rectangle[T]) ScaleF(f float32) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.ScaleF(f),
	}
}

func (r Rectangle[T]) ScaleByVector(f vector2.Float64) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.ScaleByVector(f),
	}
}

func (r Rectangle[T]) ScaleByVectorF(f vector2.Float32) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.ScaleByVectorF(f),
	}
}

func (r Rectangle[T]) ScaleByXY(x, y float64) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.ScaleByXY(x, y),
	}
}

func (r Rectangle[T]) ScaleByXYF(x, y float32) Rectangle[T] {
	return Rectangle[T]{
		position: r.position,
		size:     r.size.ScaleByXYF(x, y),
	}
}

func (r Rectangle[T]) Zoom(f float64) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.Scale(f),
		size:     r.size.Scale(f),
	}
}

func (r Rectangle[T]) ZoomF(f float32) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.ScaleF(f),
		size:     r.size.ScaleF(f),
	}
}

func (r Rectangle[T]) ZoomByVector(f vector2.Float64) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.ScaleByVector(f),
		size:     r.size.ScaleByVector(f),
	}
}

func (r Rectangle[T]) ZoomByVectorF(f vector2.Float32) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.ScaleByVectorF(f),
		size:     r.size.ScaleByVectorF(f),
	}
}

func (r Rectangle[T]) ZoomByXY(x, y float64) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.ScaleByXY(x, y),
		size:     r.size.ScaleByXY(x, y),
	}
}

func (r Rectangle[T]) ZoomByXYF(x, y float32) Rectangle[T] {
	return Rectangle[T]{
		position: r.position.ScaleByXYF(x, y),
		size:     r.size.ScaleByXYF(x, y),
	}
}

func (r Rectangle[T]) Contains(v vector2.Vector[T]) bool {
	return vector2.GreaterEq(v, r.A()) && vector2.LessEq(v, r.B())
}
