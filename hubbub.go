package hubbub

import (
	"github.com/dwimmertxt/hubbub/internal/perlin"
)

func New() *perlin.PerlinNoise {
	hb := perlin.New()
	return hb
}
