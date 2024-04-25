package vector2

import "github.com/EliCDavis/vector"

type Serializable[T vector.Number] struct {
	X T
	Y T
}

func (m Serializable[T]) Immutable() Vector[T] {
	return Vector[T]{
		x: m.X,
		y: m.Y,
	}
}
