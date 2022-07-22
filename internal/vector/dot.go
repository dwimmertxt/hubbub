package vector

func (v *Vec) Dot(ov *Vec) float64 {
	return v.Point[0] * ov.Point[0] + v.Point[1] * ov.Point[1]
}