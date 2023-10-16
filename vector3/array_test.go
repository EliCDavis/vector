package vector3_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

func TestArrayBounds(t *testing.T) {
	// ARRANGE ================================================================
	pts := []vector3.Float64{
		vector3.New(-2., 0., 0.),
		vector3.New(-2., -4., 0.),
		vector3.New(-1., -2., 1.),

		vector3.New(3., 2., 0.5),
		vector3.New(3., 1., 5.),
	}

	// ACT ====================================================================
	min, max := vector3.Array[float64](pts).Bounds()

	// ASSERT =================================================================
	assert.InDelta(t, -2, min.X(), 0.000001)
	assert.InDelta(t, -4, min.Y(), 0.000001)
	assert.InDelta(t, 0, min.Z(), 0.000001)

	assert.InDelta(t, 3, max.X(), 0.000001)
	assert.InDelta(t, 2, max.Y(), 0.000001)
	assert.InDelta(t, 5, max.Z(), 0.000001)
}

func TestArrayDistance(t *testing.T) {
	// ARRANGE ================================================================
	pts := []vector3.Float64{
		vector3.New(0., 0., 0.),
		vector3.New(0., 1., 0.),
		vector3.New(0., 1., 1.),
		vector3.New(0., 1., -1.),
	}

	// ACT ====================================================================
	dst := vector3.Array[float64](pts).Distance()

	// ASSERT =================================================================
	assert.InDelta(t, 4, dst, 0.000001)
}

func TestArrayDistanceWithOnlyOnePoint(t *testing.T) {
	// ARRANGE ================================================================
	pts := []vector3.Float64{
		vector3.New(0., 1., 0.),
	}

	// ACT ====================================================================
	dst := vector3.Array[float64](pts).Distance()

	// ASSERT =================================================================
	assert.InDelta(t, 0, dst, 0.000001)
}

func TestArrayNormalizedAndScaleAndDiv(t *testing.T) {
	// ARRANGE ================================================================
	ptCount := 1000
	pts := make([]vector3.Float64, ptCount)
	r := rand.New(rand.NewSource(42))
	for i := 0; i < ptCount; i++ {
		pts[i] = vector3.Rand(r)
	}

	// ACT ====================================================================
	dst := vector3.Array[float64](pts).
		Normalized().
		Scale(3).
		DivByConstant(2)

	// ASSERT =================================================================
	for i := 0; i < ptCount; i++ {
		assert.InDelta(t, 1.5, dst[i].Length(), 0.000001)
	}
}

func TestArrayModify(t *testing.T) {
	// ARRANGE ================================================================
	ptCount := 1000
	pts := make([]vector3.Float64, ptCount)
	r := rand.New(rand.NewSource(42))
	for i := 0; i < ptCount; i++ {
		pts[i] = vector3.Rand(r)
	}

	// ACT ====================================================================
	dst := vector3.Array[float64](pts).
		Modify(func(v vector3.Float64) vector3.Float64 {
			return v.Normalized()
		})

	// ASSERT =================================================================
	for i := 0; i < ptCount; i++ {
		assert.InDelta(t, 1, dst[i].Length(), 0.000001)
	}
}

func TestArrayStandardDeviation(t *testing.T) {
	// ARRANGE ================================================================
	ptCount := 100000
	pts := make([]vector3.Float64, ptCount)
	r := rand.New(rand.NewSource(42))

	for i := 0; i < ptCount; i++ {
		pts[i] = vector3.
			Rand(r).
			Scale(2).
			Sub(vector3.One[float64]()).
			Normalized()
	}

	// ACT ====================================================================
	average, deviation := vector3.Array[float64](pts).StandardDeviation()

	// ASSERT =================================================================
	assert.InDelta(t, 0, average.X(), 0.01)
	assert.InDelta(t, 0, average.Y(), 0.01)
	assert.InDelta(t, 0, average.Z(), 0.01)

	assert.InDelta(t, 0.5, deviation.X(), 0.1)
	assert.InDelta(t, 0.5, deviation.Y(), 0.1)
	assert.InDelta(t, 0.5, deviation.Z(), 0.1)
}

func TestArray_Add(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector3.Float64Array([]vector3.Float64{
		vector3.New(0., 0., 0.),
		vector3.New(1., 0., 0.),
		vector3.New(2., 0., 0.),
	})
	add := vector3.New(1., 2., 3.)

	// ACT ====================================================================
	added := arr.Add(add)

	// ASSERT =================================================================
	for i, v := range added {
		assert.Equal(t, arr[i].Add(add), v)
	}
}

func TestArray_Sub(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector3.Float64Array([]vector3.Float64{
		vector3.New(0., 0., 0.),
		vector3.New(1., 0., 0.),
		vector3.New(2., 0., 0.),
	})
	sub := vector3.New(1., 2., 3.)

	// ACT ====================================================================
	added := arr.Sub(sub)

	// ASSERT =================================================================
	for i, v := range added {
		assert.Equal(t, arr[i].Sub(sub), v)
	}
}

func TestArray_ContainsNaN_True(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector3.Float64Array([]vector3.Float64{
		vector3.New(0., 0., 0.),
		vector3.New(1., 0., 0.),
		vector3.New(2., math.NaN(), 0.),
	})

	// ACT ====================================================================
	assert.True(t, arr.ContainsNaN())
}

func TestArray_ContainsNaN_False(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector3.Float64Array([]vector3.Float64{
		vector3.New(0., 0., 0.),
		vector3.New(1., 0., 0.),
		vector3.New(2., 0., 0.),
	})

	// ACT ====================================================================
	assert.False(t, arr.ContainsNaN())
}

func TestArray_MaxLength(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector3.Float64Array([]vector3.Float64{
		vector3.New(0., 0., 0.),
		vector3.New(1., 0., 0.),
		vector3.New(2., 0., 0.),
	})

	// ACT ====================================================================
	assert.Equal(t, 2., arr.MaxLength())
}

func TestArray_Sum(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector3.Float64Array([]vector3.Float64{
		vector3.New(0., 1., 2.),
		vector3.New(3., 4., 5.),
		vector3.New(6., 7., 8.),
	})

	// ACT ====================================================================
	sum := arr.Sum()

	// ASSERT =================================================================
	assert.Equal(t, 9., sum.X())
	assert.Equal(t, 12., sum.Y())
	assert.Equal(t, 15., sum.Z())
}
