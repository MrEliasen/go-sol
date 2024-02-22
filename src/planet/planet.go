package planet

import (
	"github.com/MrEliasen/sol/src/config"
)

type Planet struct {
	Name    string
	Moons   []*Planet
	Primary *Planet

	Density   float64
	Radius    float64
	EarthMass float64 // in earth masses
}

func (planet *Planet) Mass() float64 {
	return planet.EarthMass * config.EARTH_MASS
}
