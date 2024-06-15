package experiments_test

import (
	"runtime"
	"testing"

	"github.com/EliCDavis/vector/vector3"
)

var addResultFloat64 vector3.Vector[float64]
var addResultMFloat64 MVector[float64]

func BenchmarkAddVector(b *testing.B) {
	var r vector3.Vector[float64]
	va := vector3.New(1., 2., 3.)
	vb := vector3.New(2., 3., 4.)

	for n := 0; n < b.N; n++ {
		r = r.Add(va).
			MultByVector(va).
			Add(va.Reciprocal()).
			MultByVector(va.Reciprocal()).
			MultByVector(vb.Reciprocal()).
			Add(vb.Reciprocal())
	}
	addResultFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddVector_NoConstants(b *testing.B) {
	var r vector3.Vector[float64]

	for n := 0; n < b.N; n++ {
		va := vector3.New(float64(n), float64(n)+1, float64(n)+2)
		vb := vector3.New(float64(n)+3, float64(n)+4, float64(n)+5)
		vaR := va.Reciprocal()
		vbR := vb.Reciprocal()
		r = r.Add(va).
			MultByVector(va).
			Add(vaR).
			MultByVector(vaR).
			MultByVector(vbR).
			Add(vbR)
	}

	addResultFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVector(b *testing.B) {
	var r MVector[float64]
	va := MVector[float64]{1., 2., 3.}
	vb := MVector[float64]{2., 3., 4.}

	for n := 0; n < b.N; n++ {
		r = r.Add(va).
			MultByVector(va).
			Add(va.Reciprocal()).
			MultByVector(va.Reciprocal()).
			MultByVector(vb.Reciprocal()).
			Add(vb.Reciprocal())
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlace(b *testing.B) {
	var r MVector[float64]
	va := MVector[float64]{1., 2., 3.}
	vb := MVector[float64]{2., 3., 4.}
	for n := 0; n < b.N; n++ {
		r.AddInPlace(va)
		r.MultInPlace(va)
		r.AddInPlace(va.Reciprocal())
		r.MultInPlace(va.Reciprocal())
		r.MultInPlace(vb.Reciprocal())
		r.AddInPlace(vb.Reciprocal())
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceUsingPointer(b *testing.B) {
	var r *MVector[float64] = &MVector[float64]{1, 2, 3}
	va := MVector[float64]{1., 2., 3.}
	vb := MVector[float64]{2., 3., 4.}

	for n := 0; n < b.N; n++ {
		r.AddInPlace(va)
		r.MultInPlace(va)
		r.AddInPlace(va.Reciprocal())
		r.MultInPlace(va.Reciprocal())
		r.MultInPlace(vb.Reciprocal())
		r.AddInPlace(vb.Reciprocal())
	}

	addResultMFloat64 = *r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceDirect(b *testing.B) {
	var r MVector[float64]
	va := MVector[float64]{1., 2., 3.}
	vb := MVector[float64]{2., 3., 4.}
	rX := 1 / vb.X
	rY := 1 / vb.Y
	rZ := 1 / vb.Z

	raX := 1 / va.X
	raY := 1 / va.Y
	raZ := 1 / va.Z
	for n := 0; n < b.N; n++ {
		r.X = (((((r.X + va.X) * va.X) + (raX)) * raX) * rX) + rX
		r.Y = (((((r.Y + va.Y) * va.Y) + (raY)) * raY) * rY) + rY
		r.Z = (((((r.Z + va.Z) * va.Z) + (raZ)) * raZ) * rZ) + rZ
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceDirect_NoConstants(b *testing.B) {
	var r MVector[float64]

	for n := 0; n < b.N; n++ {
		va := MVector[float64]{float64(n), float64(n) + 1, float64(n) + 2}
		vb := MVector[float64]{float64(n) + 3, float64(n) + 4, float64(n) + 5}

		rX := 1 / vb.X
		raX := 1 / va.X
		r.X = (((((r.X + va.X) * va.X) + (raX)) * raX) * rX) + rX

		rY := 1 / vb.Y
		raY := 1 / va.Y
		r.Y = (((((r.Y + va.Y) * va.Y) + (raY)) * raY) * rY) + rY

		rZ := 1 / vb.Z
		raZ := 1 / va.Z
		r.Z = (((((r.Z + va.Z) * va.Z) + (raZ)) * raZ) * rZ) + rZ
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceWithReturn(b *testing.B) {
	var r MVector[float64]
	va := MVector[float64]{1., 2., 3.}
	vb := MVector[float64]{2., 3., 4.}
	for n := 0; n < b.N; n++ {
		r.AddInPlaceAndReturn(va)
		r.MultInPlaceAndReturn(va)
		r.AddInPlaceAndReturn(va.Reciprocal())
		r.MultInPlaceAndReturn(va.Reciprocal())
		r.MultInPlaceAndReturn(vb.Reciprocal())
		r.AddInPlaceAndReturn(vb.Reciprocal())
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceWithReturnUsingPointer(b *testing.B) {
	var r *MVector[float64] = &MVector[float64]{1, 2, 3}
	va := MVector[float64]{1., 2., 3.}
	vb := MVector[float64]{2., 3., 4.}

	for n := 0; n < b.N; n++ {
		r.AddInPlaceAndReturn(va)
		r.MultInPlaceAndReturn(va)
		r.AddInPlaceAndReturn(va.Reciprocal())
		r.MultInPlaceAndReturn(va.Reciprocal())
		r.MultInPlaceAndReturn(vb.Reciprocal())
		r.AddInPlaceAndReturn(vb.Reciprocal())
	}

	addResultMFloat64 = *r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceTakingPointerUsingPointer(b *testing.B) {
	var r *MVector[float64] = &MVector[float64]{1, 2, 3}
	va := &MVector[float64]{1., 2., 3.}
	vb := &MVector[float64]{2., 3., 4.}

	for n := 0; n < b.N; n++ {
		ra := va.Reciprocal()
		rb := vb.Reciprocal()
		r.AddInPlaceTakingPointer(va)
		r.MultInPlaceTakingPointer(va)
		r.AddInPlaceTakingPointer(&ra)
		r.MultInPlaceTakingPointer(&ra)
		r.MultInPlaceTakingPointer(&rb)
		r.AddInPlaceTakingPointer(&rb)
	}

	addResultMFloat64 = *r
	runtime.KeepAlive(r)
	runtime.KeepAlive(addResultMFloat64)
}
