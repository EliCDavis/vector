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
		r = r.Add(va).MultByVector(va).
			Add(va.Reciprocal()).MultByVector(va.Reciprocal()).
			MultByVector(vb.Reciprocal()).Add(vb.Reciprocal())
	}
	addResultFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVector(b *testing.B) {
	var r MVector[float64]
	va := MVector[float64]{1., 2., 3.}
	vb := MVector[float64]{2., 3., 4.}

	for n := 0; n < b.N; n++ {
		r = r.Add(va).MultByVector(va).
			Add(va.Reciprocal()).MultByVector(va.Reciprocal()).
			MultByVector(vb.Reciprocal()).Add(vb.Reciprocal())
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
	for n := 0; n < b.N; n++ {
		r.X = ((r.X+va.X)*va.X+1/va.X)/va.X/vb.X + 1/vb.X
		r.Y = ((r.Y+va.Y)*va.Y+1/va.Y)/va.Y/vb.Y + 1/vb.Y
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
