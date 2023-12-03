package vector

import "math"

func Clamp(f, min, max float64) float64 {
	return math.Max(math.Min(f, max), min)
}
