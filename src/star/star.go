package star

import (
	"math"

	"github.com/MrEliasen/sol/src/config"
)

type Star struct {
	StarClass       string
	Sequence        string
	SizeSolarRadii  float64
	MassSolarMasses float64
	TemperatureK    float64
}

func (star *Star) Volume() float64 {
	return (3.0 / 4.0) * math.Pi * math.Pow(star.SizeSolarRadii, 3)
}

func (star *Star) Density() float64 {
	return star.MassSolarMasses / star.Volume()
}

func (star *Star) ThermalSpeed() float64 {
	return (2 * config.STEFAN_BOLTZMANN_CONST) * star.TemperatureK / config.PHOTON_MASS
}

func (star *Star) OuterAtmosphereDistance() float64 {
	return 5 * star.SizeSolarRadii * config.SOLAR_RADII
}

func (star *Star) Area() float64 {
	return 4 * math.Pi * math.Pow(star.SizeSolarRadii*config.SOLAR_RADII, 2)
}

func (star *Star) Luminisity() float64 {
	return (4 * math.Pi * math.Pow(star.SizeSolarRadii*config.SOLAR_RADII, 2) * config.STEFAN_BOLTZMANN_CONST * math.Pow(star.TemperatureK, 4)) / config.SOLAR_LUMINOSITY
}

func (star *Star) EscapeVelocity() (km float64) {
	return (math.Sqrt((2 * config.GRAVITATIONAL_CONST * star.MassKg()) / (star.SizeSolarRadii * config.SOLAR_RADII)))
}

func (star *Star) SolarWindVelocity() float64 {
	return star.SizeSolarRadii * config.SOLAR_WIND_VELOCITY
}

func (star *Star) SolarWindMassLossRate() float64 {
	return config.SOLAR_WIND_MASS_LOSS_DENSITY * star.Area() * star.SolarWindVelocity()
}

func (star *Star) TerminalShockDistance() float64 {
	return (star.SolarWindVelocity() / config.INTERSTELLAR_MEDIUM_DENSITY * config.APPROX_ISM_VELOCITY) * 1000 / config.AU
}

func (star *Star) IceLine() float64 {
	return 2.7 * math.Pow(star.Luminisity(), 0.5)
}

func (star *Star) MassKg() float64 {
	return star.MassSolarMasses * config.SOLAR_MASS
}

func (star *Star) HabitableZone() (start, end float64) {
	start = math.Sqrt(star.SizeSolarRadii * config.SOLAR_RADII / 2)
	end = 2 * start

	return start, end
}
