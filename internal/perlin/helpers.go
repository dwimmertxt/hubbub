package perlin

func fade(t float64) float64 {
	return ((6 * t - 15) * t + 10) * t * t * t
}

func lerp(a, b, x float64) float64 {
	return a + x * (b - a)
}