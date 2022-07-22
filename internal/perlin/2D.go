package perlin

import (
	"math"
	"hubbub/internal/vector"
)

func (pn *PerlinNoise) Get2D(x, y float64) float64 {
	// returns Perlin noise value normalised to range [0, 1]
	X 	:= int(math.Floor(x)) & 255
	Y 	:= int(math.Floor(y)) & 255
	xf 	:= x - math.Floor(x)
	yf 	:= y - math.Floor(y)

	tR 	:= vector.New(xf-1.0, yf-1.0)
	tL	:= vector.New(xf, yf-1.0)
	bR 	:= vector.New(xf-1.0, yf)
	bL 	:= vector.New(xf, yf)

	pTR := pn.perm[pn.perm[X+1]+Y+1]
	pTL := pn.perm[pn.perm[X]+Y+1]
	pBR := pn.perm[pn.perm[X+1]+Y]
	pBL := pn.perm[pn.perm[X]+Y]
		
	dTR := tR.Dot(vector.Constant(pTR))
	dTL := tL.Dot(vector.Constant(pTL))
	dBR := bR.Dot(vector.Constant(pBR))
	dBL := bL.Dot(vector.Constant(pBL))
	
	u 	:= fade(xf)
	v 	:= fade(yf)

	res := lerp(lerp(dBL, dTL, v), lerp(dBR, dTR, v), u) 

	return (res + 1) / 2
}

func (pn *PerlinNoise) Octave2D(x, y, persistence float64, octaves int) float64 {
	var (
		total 		float64 = 0.0
		frequency	float64 = 1.0
		amplitude 	float64 = 1.0
		maxValue 	float64 = 0.0
	)
	for i := 0; i < octaves; i++ {
		total += pn.Get2D(x * frequency, y * frequency) * amplitude
		maxValue += amplitude
		amplitude *= persistence
		frequency *= 2
	}
	return total / maxValue
}

