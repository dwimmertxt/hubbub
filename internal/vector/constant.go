package vector

func Constant(p int) *Vec {
	var vec *Vec
	v := p & 3
	switch v {
	case 0:
		vec = New(1.0, 1.0)
	case 1:
		vec = New(-1.0, 1.0)
	case 2:
		vec = New(-1.0, -1.0)
	case 3:
		vec = New(1.0, -1.0)
	}
	return vec
}
