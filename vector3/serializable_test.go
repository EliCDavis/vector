package vector3_test

import (
	"testing"

	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

func TestSerializable(t *testing.T) {
	s := vector3.Serializable[float64]{1, 2, 3}

	v := s.Immutable()

	assert.Equal(t, 1., v.X())
	assert.Equal(t, 2., v.Y())
	assert.Equal(t, 3., v.Z())
}
