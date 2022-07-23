package vector

import (
	"math"
)

func UnitCubePosition(point...float64) *Vec {
	var newPoint []float64
	for _, p := range point {
		newPoint = append(newPoint, math.Mod(p, 1.0))
	}
	return New(newPoint...)
}