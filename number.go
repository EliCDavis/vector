package vector

type Number interface {
	int8 | int16 | int | int32 | int64 | float32 | float64
}

type Vector interface {
	Length() float64
	LengthSquared() float64
}
