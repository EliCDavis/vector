package vector2_test

import (
	"testing"

	"github.com/EliCDavis/vector/vector2"
	"github.com/stretchr/testify/assert"
)

func TestSerializable(t *testing.T) {
	s := vector2.Serializable[float64]{1, 2}

	v := s.Immutable()

	assert.Equal(t, 1., v.X())
	assert.Equal(t, 2., v.Y())
}
