package src

import (
	"math"

	"github.com/MrEliasen/sol/src/config"
	"github.com/MrEliasen/sol/src/planet"
	"github.com/MrEliasen/sol/src/star"
	"github.com/pterm/pterm"
)

func NewStarSystem() StarSystem {
	config := star.NewClassification()
	star := config.GenerateStar()

	return StarSystem{
		Name: "Elistau",
		Star: star,
	}
}

type StarSystem struct {
	Name    string
	Star    *star.Star
	Planets map[string]*planet.Planet
}

func (system *StarSystem) RocheLimit(planet *planet.Planet) float64 {
	var a float64
	var m float64
	var M float64
	starModifier := 1.0

	if planet.Primary != nil {
		// a: avg. orbit distance to the parent/primary in AU
		a = 0.00257
		// m: the mass of the planet
		m = planet.Mass() * config.EARTH_MASS
		// M: the mass of the parent/primary
		M = planet.Primary.Mass() * config.EARTH_MASS
	} else {
		a = 2.44
		m = system.Star.Density()
		M = planet.Density
		starModifier = system.Star.SizeSolarRadii
	}

	return a * math.Pow(m/(3*M), 1.0/3.0) * starModifier
}

func (system *StarSystem) HillSphere(satelite *planet.Planet) float64 {
	var a float64
	var m float64
	var M float64

	if satelite.Primary != nil {
		// a: avg. orbit distance to the parent/primary in AU
		a = 0.00257
		// m: the mass of the planet
		m = satelite.Mass() * config.EARTH_MASS
		// M: the mass of the parent/primary
		M = satelite.Primary.Mass() * config.EARTH_MASS

		return a * math.Pow(m/(3*M), 1.0/3.0)
	} else {
		// a: avg. orbit distance to the star in AU
		a = 1.0
		// m: the mass of the satelite
		m = satelite.Mass()
		// M: the mass of the star
		M = system.Star.MassKg()
	}

	return a * math.Pow(m/(3*M), 1.0/3.0)
}

func (system *StarSystem) PlanetRocheLimit(p1, p2 *planet.Planet) (primary, satelite *planet.Planet, rocheLimit uint64) {
	primary = p1
	satelite = p2

	if p1.EarthMass < p2.EarthMass {
		primary = p2
		satelite = p1
	}

	return primary, satelite, uint64(
		2.44 * math.Pow(primary.Density/satelite.Density, 1.0/3.0) * primary.Radius,
	)
}

func (system *StarSystem) Display() {
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
		{"Class", "Sequence", "Size (R☉)", "Mass (M☉)", "Density (g/cm)", "Solar Wind Speed (km/s)", "Escape Velocity (km/s)", "Termal Shock Distance", "Ice Line", "Temperature (c / k)", "Luminosity"},
		{
			system.Star.StarClass,
			system.Star.Sequence,
			pterm.Sprintf("%f ", system.Star.SizeSolarRadii),
			pterm.Sprintf("%f", system.Star.MassSolarMasses),
			pterm.Sprintf("%f", system.Star.Density()),
			pterm.Sprintf("%f", system.Star.SolarWindVelocity()/1000),
			pterm.Sprintf("%f", system.Star.EscapeVelocity()/1000),
			pterm.Sprintf("%.5f", system.Star.TerminalShockDistance()),
			pterm.Sprintf("%.5f", system.Star.IceLine()),
			pterm.Sprintf("%f / %f", system.Star.TemperatureK, config.KelvinToCelcius(system.Star.TemperatureK)),
			pterm.Sprintf("%.5f", system.Star.Luminisity()),
			"",
		},
	}).Render()
}
