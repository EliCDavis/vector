package experiments_test

import (
	"math"
	"runtime"
	"testing"
)

var minResult float64
var minResult32 float32
var minResultInt int

func BenchmarkMinBuiltinFloat64(b *testing.B) {
	var r float64

	for n := 0; n < b.N; n++ {
		r += min(float64(n)-1, float64(n))
	}

	minResult = r
	runtime.KeepAlive(r)
}

func BenchmarkMinMathPackageFloat64(b *testing.B) {
	var r float64

	for n := 0; n < b.N; n++ {
		r += math.Min(float64(n)-1, float64(n))
	}

	minResult = r
	runtime.KeepAlive(r)
}

func BenchmarkMinBuiltinFloat32(b *testing.B) {
	var r float32

	for n := 0; n < b.N; n++ {
		r += min(float32(n)-1, float32(n))
	}

	minResult32 = r
	runtime.KeepAlive(r)
}

func BenchmarkMinMathPackageFloat32(b *testing.B) {
	var r float32

	for n := 0; n < b.N; n++ {
		r += float32(math.Min(float64(n)-1, float64(n)))
	}

	minResult32 = r
	runtime.KeepAlive(r)
}

func BenchmarkMinBuiltinInt(b *testing.B) {
	var r int

	for n := 0; n < b.N; n++ {
		r += min(int(n)-1, int(n))
	}

	minResultInt = r
	runtime.KeepAlive(r)
}

func BenchmarkMinMathPackageInt(b *testing.B) {
	var r int

	for n := 0; n < b.N; n++ {
		r += int(math.Min(float64(n)-1, float64(n)))
	}

	minResultInt = r
	runtime.KeepAlive(r)
}
