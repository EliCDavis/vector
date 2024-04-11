package vector

func Clamp(f, vmin, vmax float64) float64 {
	return max(min(f, vmax), vmin)
}
