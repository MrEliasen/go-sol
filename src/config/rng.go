package config

import "math/rand"

var (
	seed int64
	rng  *rand.Rand
)

func SeedRng(s int64) {
	seed = s
}

func GetRng() *rand.Rand {
	if rng != nil {
		return rng
	}

	rng = rand.New(rand.NewSource(seed))
	return rng
}
