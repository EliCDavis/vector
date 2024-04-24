package vector3

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/EliCDavis/vector"
)

const componentCount = 3

func (v Vector[T]) Write(out io.Writer, endian binary.ByteOrder) (err error) {
	switch vv := any(v).(type) {
	case Float64:
		bytes := make([]byte, 8*componentCount)
		endian.PutUint64(bytes, math.Float64bits(vv.x))
		endian.PutUint64(bytes[8:], math.Float64bits(vv.y))
		endian.PutUint64(bytes[16:], math.Float64bits(vv.z))
		_, err = out.Write(bytes)
		return
	case Float32:
		bytes := make([]byte, 4*componentCount)
		endian.PutUint32(bytes, math.Float32bits(vv.x))
		endian.PutUint32(bytes[4:], math.Float32bits(vv.y))
		endian.PutUint32(bytes[8:], math.Float32bits(vv.z))
		_, err = out.Write(bytes)
		return

	case Int8:
		_, err = out.Write([]byte{
			byte(vv.x),
			byte(vv.y),
			byte(vv.z),
		})
		return

	case Int16:
		bytes := make([]byte, 2*componentCount)
		endian.PutUint16(bytes, uint16(vv.x))
		endian.PutUint16(bytes[2:], uint16(vv.y))
		endian.PutUint16(bytes[4:], uint16(vv.z))
		_, err = out.Write(bytes)
		return

	case Int32:
		bytes := make([]byte, 4*componentCount)
		endian.PutUint32(bytes, uint32(vv.x))
		endian.PutUint32(bytes[4:], uint32(vv.y))
		endian.PutUint32(bytes[8:], uint32(vv.z))
		_, err = out.Write(bytes)
		return

	case Int64:
		bytes := make([]byte, 8*componentCount)
		endian.PutUint64(bytes, uint64(vv.x))
		endian.PutUint64(bytes[8:], uint64(vv.y))
		endian.PutUint64(bytes[16:], uint64(vv.z))
		_, err = out.Write(bytes)
		return
	}

	panic(fmt.Errorf("write unimplemented type: %#v", v))
}

func Read[T vector.Number](in io.Reader, endian binary.ByteOrder) (v Vector[T], err error) {
	switch any(v).(type) {
	case Float64:
		vv, err := ReadFloat64(in, endian)
		return any(vv).(Vector[T]), err

	case Float32:
		vv, err := ReadFloat32(in, endian)
		return any(vv).(Vector[T]), err

	case Int8:
		vv, err := ReadInt8(in)
		return any(vv).(Vector[T]), err

	case Int16:
		vv, err := ReadInt16(in, endian)
		return any(vv).(Vector[T]), err

	case Int32:
		vv, err := ReadInt32(in, endian)
		return any(vv).(Vector[T]), err

	case Int64:
		vv, err := ReadInt64(in, endian)
		return any(vv).(Vector[T]), err
	}

	panic(fmt.Errorf("read unimplemented type: %#v", v))
}

func ReadFloat64(in io.Reader, endian binary.ByteOrder) (Vector[float64], error) {
	buf := make([]byte, componentCount*8)
	_, err := io.ReadFull(in, buf)
	return Vector[float64]{
		x: math.Float64frombits(endian.Uint64(buf)),
		y: math.Float64frombits(endian.Uint64(buf[8:])),
		z: math.Float64frombits(endian.Uint64(buf[16:])),
	}, err
}

func ReadFloat32(in io.Reader, endian binary.ByteOrder) (Vector[float32], error) {
	buf := make([]byte, componentCount*4)
	_, err := io.ReadFull(in, buf)
	return Vector[float32]{
		x: math.Float32frombits(endian.Uint32(buf)),
		y: math.Float32frombits(endian.Uint32(buf[4:])),
		z: math.Float32frombits(endian.Uint32(buf[8:])),
	}, err
}

func ReadInt8(in io.Reader) (Vector[int8], error) {
	buf := make([]byte, componentCount)
	_, err := io.ReadFull(in, buf)
	return Vector[int8]{
		x: int8(buf[0]),
		y: int8(buf[1]),
		z: int8(buf[2]),
	}, err
}

func ReadInt16(in io.Reader, endian binary.ByteOrder) (Vector[int16], error) {
	buf := make([]byte, componentCount*2)
	_, err := io.ReadFull(in, buf)
	return Vector[int16]{
		x: int16(endian.Uint16(buf)),
		y: int16(endian.Uint16(buf[2:])),
		z: int16(endian.Uint16(buf[4:])),
	}, err
}

func ReadInt32(in io.Reader, endian binary.ByteOrder) (Vector[int32], error) {
	buf := make([]byte, componentCount*4)
	_, err := io.ReadFull(in, buf)
	return Vector[int32]{
		x: int32(endian.Uint32(buf)),
		y: int32(endian.Uint32(buf[4:])),
		z: int32(endian.Uint32(buf[8:])),
	}, err
}

func ReadInt64(in io.Reader, endian binary.ByteOrder) (Vector[int64], error) {
	buf := make([]byte, componentCount*8)
	_, err := io.ReadFull(in, buf)
	return Vector[int64]{
		x: int64(endian.Uint64(buf)),
		y: int64(endian.Uint64(buf[8:])),
		z: int64(endian.Uint64(buf[16:])),
	}, err
}
