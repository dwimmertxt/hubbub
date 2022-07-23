package hubbub

import (
	"hubbub/internal/perlin"
)

func New() *perlin.PerlinNoise {
	hb := perlin.New()
	return hb
}
