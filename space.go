package vector

type Space[T any] interface {
	Distance(a, b T) float64
	Add(a, b T) T
	Sub(a, b T) T
	Scale(a T, amount float64) T
	Dot(a, b T) float64
	Length(a T) float64
	Normalized(a T) T
	Lerp(a, b T, time float64) T
}
