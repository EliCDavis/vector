package vector

import "math"

func clamp(f, min, max float64) float64 {
	return math.Max(math.Min(f, max), min)
}
