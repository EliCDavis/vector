package vector4

import "github.com/EliCDavis/vector"

type Serializable[T vector.Number] struct {
	X T
	Y T
	Z T
	W T
}

func (m Serializable[T]) Immutable() Vector[T] {
	return Vector[T]{
		x: m.X,
		y: m.Y,
		z: m.Z,
		w: m.W,
	}
}
