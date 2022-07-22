package perlin

type PerlinNoise struct {
	perm 	[512]int
}

func New() *PerlinNoise {
	pn 		   := new(PerlinNoise)
	pn.perm 	= permutationTable()
	return pn
}

