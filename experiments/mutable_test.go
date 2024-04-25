package experiments_test

import "github.com/EliCDavis/vector"

type MVector[T vector.Number] struct {
	X T
	Y T
	Z T
}

func (v MVector[T]) Add(a MVector[T]) MVector[T] {
	return MVector[T]{
		X: v.X + a.X,
		Y: v.Y + a.Y,
		Z: v.Z + a.Z,
	}
}

func (v *MVector[T]) AddInPlace(a MVector[T]) {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
}

func (v *MVector[T]) AddInPlaceAndReturn(a MVector[T]) MVector[T] {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
	return *v
}

func (v *MVector[T]) AddInPlaceTakingPointer(a *MVector[T]) {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
}
