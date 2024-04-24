package vector3

import "github.com/EliCDavis/vector"

type Mutable[T vector.Number] struct {
	X T
	Y T
	Z T
}

type (
	MutFloat64 = Mutable[float64]
	MutFloat32 = Mutable[float32]
	MutInt     = Mutable[int]
	MutInt64   = Mutable[int64]
	MutInt32   = Mutable[int32]
	MutInt16   = Mutable[int16]
	MutInt8    = Mutable[int8]
)
