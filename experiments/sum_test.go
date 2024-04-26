package experiments_test

import (
	"fmt"
	"testing"

	"github.com/EliCDavis/vector/vector3"
)

var arrLenToTest = []int{
	100,
	1_000,
	10_000,
	100_000,
	1_000_000,
	10_000_000,
}

var resultFloat64 []vector3.Float64
var resultMFloat64 []MVector[float64]

func BenchmarkSumVector(b *testing.B) {
	var vals []vector3.Float64
	for _, testLen := range arrLenToTest {
		vals = make([]vector3.Float64, testLen)
		add := vector3.New(1., 2., 3.)

		b.Run(fmt.Sprintf("Len-%d", testLen), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i, v := range vals {
					vals[i] = v.Add(add)
				}
			}
		})
	}
	resultFloat64 = vals
}

func BenchmarkSumMutableVector(b *testing.B) {
	var vals []MVector[float64]
	for _, testLen := range arrLenToTest {
		vals = make([]MVector[float64], testLen)
		add := MVector[float64]{1., 2., 3.}

		b.Run(fmt.Sprintf("Len-%d", testLen), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i, v := range vals {
					vals[i] = MVector[float64]{
						X: v.X + add.X,
						Y: v.Y + add.Y,
						Z: v.Z + add.Z,
					}
				}
			}
		})
	}
	resultMFloat64 = vals
}
