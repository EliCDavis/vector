package rect2

import (
	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector2"
)

type Rectangle[T vector.Number] struct {
	xy vector2.Vector[T]
	wh vector2.Vector[T]
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
		xy: xy,
		wh: wh,
	}
}

func Zero[T vector.Number]() Rectangle[T] {
	return Rectangle[T]{
		xy: vector2.Zero[T](),
		wh: vector2.Zero[T](),
	}
}

func One[T vector.Number]() Rectangle[T] {
	return Rectangle[T]{
		xy: vector2.Zero[T](),
		wh: vector2.One[T](),
	}
}

func (r Rectangle[T]) A() vector2.Vector[T] {
	return r.xy
}

func (r Rectangle[T]) SetA(a vector2.Vector[T]) Rectangle[T] {
	dxy := a.Sub(r.xy)
	return Rectangle[T]{
		xy: a,
		wh: r.wh.Sub(dxy),
	}
}

func (r Rectangle[T]) B() vector2.Vector[T] {
	return r.xy.Add(r.wh)
}

func (r Rectangle[T]) SetB(b vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: b,
	}
}

func (r Rectangle[T]) HorizontalLine(y T) (vector2.Vector[T], vector2.Vector[T]) {
	return vector2.New(r.xy.X(), y), vector2.New(r.xy.X()+r.wh.X(), y)
}

func (r Rectangle[T]) VerticalLine(x T) (vector2.Vector[T], vector2.Vector[T]) {
	return vector2.New(x, r.xy.Y()), vector2.New(x, r.xy.Y()+r.wh.Y())
}

func (r Rectangle[T]) Center() vector2.Vector[T] {
	return r.xy.Add(r.wh.ScaleF(0.5))
}

func (v Rectangle[T]) ToFloat64() Rectangle[float64] {
	return Rectangle[float64]{
		xy: v.xy.ToFloat64(),
		wh: v.wh.ToFloat64(),
	}
}

func (v Rectangle[T]) ToFloat32() Rectangle[float32] {
	return Rectangle[float32]{
		xy: v.xy.ToFloat32(),
		wh: v.wh.ToFloat32(),
	}
}

func (v Rectangle[T]) ToInt() Rectangle[int] {
	return Rectangle[int]{
		xy: v.xy.ToInt(),
		wh: v.wh.ToInt(),
	}
}

func (v Rectangle[T]) ToInt32() Rectangle[int32] {
	return Rectangle[int32]{
		xy: v.xy.ToInt32(),
		wh: v.wh.ToInt32(),
	}
}

func (v Rectangle[T]) ToInt64() Rectangle[int64] {
	return Rectangle[int64]{
		xy: v.xy.ToInt64(),
		wh: v.wh.ToInt64(),
	}
}

// X returns the x of the xy component
func (r Rectangle[T]) X() T {
	return r.xy.X()
}

// SetX changes the x of the xy component of the rectangle
func (r Rectangle[T]) SetX(newX T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.SetX(newX),
		wh: r.wh,
	}
}

func (r Rectangle[T]) Dx(dX T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.Dx(dX),
		wh: r.wh,
	}
}

// Y returns the y of the xy component
func (r Rectangle[T]) Y() T {
	return r.xy.Y()
}

// SetY changes the y of the xy component of the rectangle
func (r Rectangle[T]) SetY(newY T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.SetY(newY),
		wh: r.wh,
	}
}

func (r Rectangle[T]) Dy(dY T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.Dy(dY),
		wh: r.wh,
	}
}

// Width returns the x of the wh component
func (r Rectangle[T]) Width() T {
	return r.wh.X()
}

// SetWidth changes the x of the wh component of the rectangle
func (r Rectangle[T]) SetWidth(newW T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.SetX(newW),
	}
}

func (r Rectangle[T]) Dw(dW T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.Dx(dW),
	}
}

// Y returns the y of the wh component
func (r Rectangle[T]) Height() T {
	return r.wh.Y()
}

// SetHeight changes the y of the wh component of the rectangle
func (r Rectangle[T]) SetHeight(newH T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.SetY(newH),
	}
}

func (r Rectangle[T]) Dh(dH T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.Dy(dH),
	}
}

// XY returns the xy component
func (r Rectangle[T]) Position() vector2.Vector[T] {
	return r.xy
}

// SetXY changes the xy component of the rectangle
func (r Rectangle[T]) SetXY(newXY vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		xy: newXY,
		wh: r.wh,
	}
}

func (r Rectangle[T]) Dxy(dXY vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.Add(dXY),
		wh: r.wh,
	}
}

// WH returns the xy component
func (r Rectangle[T]) WH() vector2.Vector[T] {
	return r.wh
}

// SetXY changes the wh component of the rectangle
func (r Rectangle[T]) SetWH(newWH vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: newWH,
	}
}

func (r Rectangle[T]) Dwh(dWH vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.Add(dWH),
	}
}

// Round takes each component of the rectangle and rounds it to the nearest whole
// number
func (v Rectangle[T]) Round() Rectangle[T] {
	return New(
		v.xy.Round(),
		v.wh.Round(),
	)
}

// RoundToInt takes each component of the rectangle and rounds it to the nearest
// whole number, and then casts it to a int
func (v Rectangle[T]) RoundToInt() Rectangle[int] {
	return New(
		v.xy.RoundToInt(),
		v.wh.RoundToInt(),
	)
}

// Ceil applies the ceil math operation to each component of the rectangle
func (v Rectangle[T]) Ceil() Rectangle[T] {
	return New(
		v.xy.Ceil(),
		v.wh.Ceil(),
	)
}

// CeilToInt applies the ceil math operation to each component of the rectangle,
// and then casts it to a int
func (v Rectangle[T]) CeilToInt() Rectangle[int] {
	return New(
		v.xy.CeilToInt(),
		v.wh.CeilToInt(),
	)
}

// Floor applies the floor math operation to each component of the rectangle
func (v Rectangle[T]) Floor() Rectangle[T] {
	return New(
		v.xy.Floor(),
		v.wh.Floor(),
	)
}

// FloorToInt applies the floor math operation to each component of the rectangle,
// and then casts it to a int
func (v Rectangle[T]) FloorToInt() Rectangle[int] {
	return New(
		v.xy.FloorToInt(),
		v.wh.FloorToInt(),
	)
}


func (r Rectangle[T]) ShiftXY(x, y T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.AddXY(x, y),
		wh: r.wh,
	}
}

func (r Rectangle[T]) Delta(xy vector2.Vector[T], wh vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.Add(xy),
		wh: r.wh.Add(wh),
	}
}

func (r Rectangle[T]) DeltaXYWH(x, y, w, h T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.AddXY(x, y),
		wh: r.wh.AddXY(w, h),
	}
}

func (r Rectangle[T]) ShrinkXYWH(left, top, right, bottom T) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.AddXY(left, top),
		wh: r.wh.AddXY(-left-right, -top-bottom),
	}
}

func (r Rectangle[T]) Scale(f float64) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.Scale(f),
	}
}

func (r Rectangle[T]) ScaleF(f float32) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.ScaleF(f),
	}
}

func (r Rectangle[T]) ScaleByVector(f vector2.Float64) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.ScaleByVector(f),
	}
}

func (r Rectangle[T]) ScaleByVectorF(f vector2.Float32) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.ScaleByVectorF(f),
	}
}

func (r Rectangle[T]) ScaleByXYF(x, y float32) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy,
		wh: r.wh.ScaleByXYF(x, y),
	}
}

func (r Rectangle[T]) Zoom(f float64) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.Scale(f),
		wh: r.wh.Scale(f),
	}
}

func (r Rectangle[T]) ZoomF(f float32) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.ScaleF(f),
		wh: r.wh.ScaleF(f),
	}
}

func (r Rectangle[T]) ZoomByVector(f vector2.Float64) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.ScaleByVector(f),
		wh: r.wh.ScaleByVector(f),
	}
}

func (r Rectangle[T]) ZoomByVectorF(f vector2.Float32) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.ScaleByVectorF(f),
		wh: r.wh.ScaleByVectorF(f),
	}
}

func (r Rectangle[T]) ZoomByXYF(x, y float32) Rectangle[T] {
	return Rectangle[T]{
		xy: r.xy.ScaleByXYF(x, y),
		wh: r.wh.ScaleByXYF(x, y),
	}
}

func (r Rectangle[T]) Contains(v vector2.Vector[T]) bool {
	return vector2.GreaterEq(v, r.A()) && vector2.LessEq(v, r.B())
}
