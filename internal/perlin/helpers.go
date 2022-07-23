package perlin

func fade(t float64) float64 {
	return ((6.0 * t - 15.0) * t + 10.0) * t * t * t
}

func lerp(a, b, x float64) float64 {
	return a + x * (b - a)
}