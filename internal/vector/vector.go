package vector

type Vec struct {
	Point 	[]float64
}

func New(p...float64) *Vec {
	v := new(Vec)
	v.setPoint(p...)
	return v
}

func (v *Vec) setPoint(p ...float64) {
	v.Point = p
}