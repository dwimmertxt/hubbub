package perlin

import (
	"math/rand"
)

func permutationTable() [512]int {
	var perm [512]int
	for i := 0; i < 256; i++ {
		perm[i] = i
	}
	rand.Shuffle(len(perm), func(i, j int) {
		perm[i], perm[j] = perm[j], perm[i]
	})
	for i := 256; i < 512; i++ {
		perm[i] = perm[i- 256]
	}
	return perm
}