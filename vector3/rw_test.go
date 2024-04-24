package vector3_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

type testCaseI interface {
	test(t *testing.T)
}

type readWriteTestCase[T vector.Number] struct {
	val vector3.Vector[T]
}

func (tc readWriteTestCase[T]) test(t *testing.T) {
	buf := &bytes.Buffer{}

	var v T

	assert.NoError(t, tc.val.Write(buf, binary.LittleEndian))
	assert.Equal(t, binary.Size(v)*3, buf.Len())
	back, err := vector3.Read[T](buf, binary.LittleEndian)
	assert.NoError(t, err)
	assert.Equal(t, tc.val, back)
}

func TestReadWrite(t *testing.T) {
	tests := map[string]testCaseI{
		"float64": readWriteTestCase[float64]{
			val: vector3.New(1., 2., 3.),
		},
		"float32": readWriteTestCase[float32]{
			val: vector3.New[float32](1., 2., 3.),
		},
		"int8": readWriteTestCase[int8]{
			val: vector3.New[int8](1., 2., 3.),
		},
		"int16": readWriteTestCase[int16]{
			val: vector3.New[int16](1., 2., 3.),
		},
		"int32": readWriteTestCase[int32]{
			val: vector3.New[int32](1., 2., 3.),
		},
		"int64": readWriteTestCase[int64]{
			val: vector3.New[int64](1., 2., 3.),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.test(t)
		})
	}
}
