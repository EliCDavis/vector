package vector3_test

import (
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

func TestArrayNormalized(t *testing.T) {
	// ARRANGE ================================================================
	ptCount := 1000
	pts := make([]vector3.Float64, ptCount)
	for i := 0; i < ptCount; i++ {
		pts[i] = vector3.Rand()
	}

	// ACT ====================================================================
	dst := vector3.Array[float64](pts).Normalized()

	// ASSERT =================================================================
	for i := 0; i < ptCount; i++ {
		assert.InDelta(t, 1., dst[i].Length(), 0.000001)
	}
}
