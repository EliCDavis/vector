package vector4_test

import (
	"testing"

	"github.com/EliCDavis/vector/vector4"
	"github.com/stretchr/testify/assert"
)

func TestSerializable(t *testing.T) {
	s := vector4.Serializable[float64]{1, 2, 3, 4}

	v := s.Immutable()

	assert.Equal(t, 1., v.X())
	assert.Equal(t, 2., v.Y())
	assert.Equal(t, 3., v.Z())
	assert.Equal(t, 4., v.W())
}
