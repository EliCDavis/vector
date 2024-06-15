package vector2_test

import (
	"math/rand"
	"testing"

	"github.com/EliCDavis/vector/vector2"
	"github.com/stretchr/testify/assert"
)

func TestArrayBounds(t *testing.T) {
	// ARRANGE ================================================================
	pts := []vector2.Float64{
		vector2.New(-2., 0.),
		vector2.New(-2., -4.),
		vector2.New(-1., -2.),

		vector2.New(3., 2.),
		vector2.New(3., 1.),
	}

	// ACT ====================================================================
	min, max := vector2.Array[float64](pts).Bounds()

	// ASSERT =================================================================
	assert.InDelta(t, -2, min.X(), 0.000001)
	assert.InDelta(t, -4, min.Y(), 0.000001)

	assert.InDelta(t, 3, max.X(), 0.000001)
	assert.InDelta(t, 2, max.Y(), 0.000001)
}

func TestArrayBounds_PanicsOnZeroPoints(t *testing.T) {
	// ARRANGE ================================================================
	pts := []vector2.Float64{}

	// ACT ====================================================================
	assert.PanicsWithError(t, "can not compute bounds from 0 vector elements", func() {
		vector2.Float64Array(pts).Bounds()
	})

}

func TestArrayNormalizedAndScaleAndDiv(t *testing.T) {
	// ARRANGE ================================================================
	ptCount := 1000
	pts := make([]vector2.Float64, ptCount)
	r := rand.New(rand.NewSource(42))
	for i := 0; i < ptCount; i++ {
		pts[i] = vector2.Rand(r)
	}

	// ACT ====================================================================
	dst := vector2.Array[float64](pts).
		Normalized().
		Scale(3)

	// ASSERT =================================================================
	for i := 0; i < ptCount; i++ {
		assert.InDelta(t, 3, dst[i].Length(), 0.000001)
	}
}

func TestArray_Add(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector2.Float64Array([]vector2.Float64{
		vector2.New(0., 0.),
		vector2.New(1., 0.),
		vector2.New(2., 0.),
	})
	add := vector2.New(1., 2.)

	// ACT ====================================================================
	added := arr.Add(add)

	// ASSERT =================================================================
	for i, v := range added {
		assert.Equal(t, arr[i].Add(add), v)
	}
}

func TestArray_AddInplace(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector2.Float64Array([]vector2.Float64{
		vector2.New(0., 0.),
		vector2.New(1., 0.),
		vector2.New(2., 0.),
	})
	add := vector2.New(1., 2.)

	// ACT ====================================================================
	arr.AddInplace(add)

	// ASSERT =================================================================
	assert.Equal(t, vector2.New(1., 2.), arr[0])
	assert.Equal(t, vector2.New(2., 2.), arr[1])
	assert.Equal(t, vector2.New(3., 2.), arr[2])
}

func TestArray_ScaleInplace(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector2.Float64Array([]vector2.Float64{
		vector2.New(0., 0.),
		vector2.New(1., 0.),
		vector2.New(2., 0.),
	})

	// ACT ====================================================================
	arr.ScaleInplace(2)

	// ASSERT =================================================================
	assert.Equal(t, vector2.New(0., 0.), arr[0])
	assert.Equal(t, vector2.New(2., 0.), arr[1])
	assert.Equal(t, vector2.New(4., 0.), arr[2])
}

func TestArray_Sum(t *testing.T) {
	// ARRANGE ================================================================
	arr := vector2.Float64Array([]vector2.Float64{
		vector2.New(0., 1.),
		vector2.New(3., 4.),
		vector2.New(6., 7.),
	})

	// ACT ====================================================================
	sum := arr.Sum()

	// ASSERT =================================================================
	assert.Equal(t, 9., sum.X())
	assert.Equal(t, 12., sum.Y())
}
