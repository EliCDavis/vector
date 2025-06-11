package vector2

import "github.com/EliCDavis/vector"

type Space[T vector.Number] struct{}

func (Space[T]) Distance(a, b Vector[T]) float64 {
	return a.Distance(b)
}

func (Space[T]) Add(a, b Vector[T]) Vector[T] {
	return a.Add(b)
}

func (Space[T]) Sub(a, b Vector[T]) Vector[T] {
	return a.Sub(b)
}

func (Space[T]) Scale(a Vector[T], amount float64) Vector[T] {
	return a.Scale(amount)
}

func (Space[T]) Dot(a, b Vector[T]) float64 {
	return a.Dot(b)
}

func (Space[T]) Length(a Vector[T]) float64 {
	return a.Length()
}

func (Space[T]) Normalized(a Vector[T]) Vector[T] {
	return a.Normalized()
}

func (Space[T]) Lerp(a, b Vector[T], time float64) Vector[T] {
	return Lerp(a, b, time)
}
