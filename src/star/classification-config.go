package star

import (
	"github.com/MrEliasen/sol/src/config"
)

type StarConfig struct {
	Name         string
	SizeMin      float64
	SizeMax      float64
	MassMin      float64
	MassMax      float64
	TempRangeMin float64
	TempRangeMax float64
}

func (starConfig *StarConfig) GenerateStar() *Star {
	rng := config.GetRng()

	temp := starConfig.TempRangeMin + rng.Float64()*(starConfig.TempRangeMax-starConfig.TempRangeMin)

	star := &Star{
		StarClass:       starConfig.Name,
		SizeSolarRadii:  starConfig.SizeMin + rng.Float64()*(starConfig.SizeMax-starConfig.SizeMin),
		MassSolarMasses: starConfig.MassMin + rng.Float64()*(starConfig.MassMax-starConfig.MassMin),
		TemperatureK:    temp,
	}

	star.Sequence = calculateSpectralClass(star)

	return star
}

func NewClassification() StarConfig {
	rng := config.GetRng()
	num := rng.Float64() * 100

	switch true {
	case num <= 3.0e-005:
		return OClass
	case num <= 0.13:
		return BClass
	case num <= 0.6:
		return AClass
	case num <= 3.0:
		return FClass
	case num <= 7.6:
		return GClass
	case num <= 12.1:
		return KClass
	default:
		return MClass
	}
}

func calculateSpectralClass(star *Star) string {
	solarLumen := star.Luminisity()

	for _, v := range spectralClassess[star.StarClass] {
		if solarLumen < v.LuminosityMin || solarLumen > v.LuminosityMax {
			continue
		}

		if v.TempKelvinMin != 0.0 && v.TempKelvinMax != 0.0 {
			if star.TemperatureK < v.TempKelvinMin || star.TemperatureK > v.TempKelvinMax {
				continue
			}
		}

		return v.Name
	}

	return "Unknown"
}

var (
	MClass = StarConfig{
		Name:         "M",
		SizeMin:      0.1,
		SizeMax:      0.7,
		MassMin:      0.1,
		MassMax:      0.6,
		TempRangeMin: 0,
		TempRangeMax: 3499,
	}
	KClass = StarConfig{
		Name:         "K",
		SizeMin:      0.8,
		SizeMax:      0.9,
		MassMin:      0.6,
		MassMax:      0.8,
		TempRangeMin: 3500,
		TempRangeMax: 4999,
	}
	GClass = StarConfig{
		Name:         "G",
		SizeMin:      0.9,
		SizeMax:      1.1,
		MassMin:      0.8,
		MassMax:      1.2,
		TempRangeMin: 5000,
		TempRangeMax: 5999,
	}
	FClass = StarConfig{
		Name:         "F",
		SizeMin:      1.1,
		SizeMax:      1.4,
		MassMin:      1.2,
		MassMax:      1.4,
		TempRangeMin: 6000,
		TempRangeMax: 7499,
	}
	AClass = StarConfig{
		Name:         "A",
		SizeMin:      1.4,
		SizeMax:      1.8,
		MassMin:      1.4,
		MassMax:      2.0,
		TempRangeMin: 7500,
		TempRangeMax: 9999,
	}
	BClass = StarConfig{
		Name:         "B",
		SizeMin:      1.8,
		SizeMax:      6.6,
		MassMin:      2.0,
		MassMax:      15.0,
		TempRangeMin: 10000,
		TempRangeMax: 29999,
	}
	OClass = StarConfig{
		Name:         "O",
		SizeMin:      6.6,
		SizeMax:      1700.0,
		MassMin:      15.0,
		MassMax:      30.0,
		TempRangeMin: 30000,
		TempRangeMax: 60000,
	}
)
