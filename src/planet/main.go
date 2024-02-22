package planet

import "github.com/MrEliasen/sol/src/config"

func NewPlanet() Planet {
	rng := config.GetRng()
	num := rng.Float64() * 100

	switch true {
	case num <= 100:
		return Planet{}
	default:
		return Planet{}
	}
}
