package config

const (
	AU                           = 149.6e6       // in km
	EARTH_MASS                   = 5.97e24       // in kg
	SOLAR_MASS                   = 1.99e30       // in kg
	SOLAR_RADII                  = 6.96e8        // in meters
	SOLAR_WIND_MASS_LOSS_DENSITY = 10e-12        // kg/m3
	SOLAR_WIND_VELOCITY          = 400e03        // m/s
	SOLAR_ESCAPE_VELOCITY        = 617.54        // km/s
	SOLAR_LUMINOSITY             = 3.828e26      // solar lumens
	GRAVITATIONAL_CONST          = 6.674e-11     // :)
	INTERSTELLAR_MEDIUM_DENSITY  = 1             // atoms/cm^3
	APPROX_ISM_VELOCITY          = 25            // km/s
	STEFAN_BOLTZMANN_CONST       = 5.67e-8       // W/(m^2 K^4)
	PHOTON_MASS                  = 1.6726219e-27 // in kg
)

func CelciusToKelvin(c float64) float64 {
	return c + 273.15
}

func KelvinToCelcius(k float64) float64 {
	return k - 273.15
}
