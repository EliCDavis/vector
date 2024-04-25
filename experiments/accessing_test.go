package experiments_test

import (
	"runtime"
	"testing"

	"github.com/EliCDavis/vector/vector3"
)

var accessResult int

func BenchmarkAccessVector(b *testing.B) {
	var r int

	for n := 0; n < b.N; n++ {
		a := vector3.New(n, n, n)
		r = a.X()
	}

	accessResult = r
	runtime.KeepAlive(r)
}

func BenchmarkAccessMutableVector(b *testing.B) {
	var r int

	for n := 0; n < b.N; n++ {
		a := MVector[int]{n, n, n}
		r = a.X
	}

	accessResult = r
	runtime.KeepAlive(r)
}
