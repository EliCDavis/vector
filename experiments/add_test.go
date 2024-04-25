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
	a := vector3.New(1., 2., 3.)

	for n := 0; n < b.N; n++ {
		r = r.Add(a)
	}
	addResultFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVector(b *testing.B) {
	var r MVector[float64]
	a := MVector[float64]{1., 2., 3.}

	for n := 0; n < b.N; n++ {
		r = r.Add(a)
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlace(b *testing.B) {
	var r MVector[float64]
	a := MVector[float64]{1., 2., 3.}
	for n := 0; n < b.N; n++ {
		r.AddInPlace(a)
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceUsingPointer(b *testing.B) {
	var r *MVector[float64] = &MVector[float64]{1, 2, 3}
	a := MVector[float64]{1., 2., 3.}

	for n := 0; n < b.N; n++ {
		r.AddInPlace(a)
	}

	addResultMFloat64 = *r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceWithReturn(b *testing.B) {
	var r MVector[float64]
	a := MVector[float64]{1., 2., 3.}
	for n := 0; n < b.N; n++ {
		r.AddInPlaceAndReturn(a)
	}
	addResultMFloat64 = r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceWithReturnUsingPointer(b *testing.B) {
	var r *MVector[float64] = &MVector[float64]{1, 2, 3}
	a := MVector[float64]{1., 2., 3.}

	for n := 0; n < b.N; n++ {
		r.AddInPlaceAndReturn(a)
	}

	addResultMFloat64 = *r
	runtime.KeepAlive(r)
}

func BenchmarkAddMutableVectorInPlaceTakingPointerUsingPointer(b *testing.B) {
	var r *MVector[float64] = &MVector[float64]{1, 2, 3}
	a := &MVector[float64]{1., 2., 3.}

	for n := 0; n < b.N; n++ {
		r.AddInPlaceTakingPointer(a)
	}

	addResultMFloat64 = *r
	runtime.KeepAlive(r)
}
