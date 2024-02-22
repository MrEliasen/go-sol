package star

type SpectralClassConfiguration struct {
	Name          string
	LuminosityMax float64
	LuminosityMin float64
	TempKelvinMax float64
	TempKelvinMin float64
}

var spectralClassess = map[string][]SpectralClassConfiguration{
	"O": {
		{
			Name:          "Ia/Ib (Supergiant)",
			LuminosityMax: 999999,
			LuminosityMin: 7500,
		},
		{
			Name:          "II (Bright Giant)",
			LuminosityMax: 7499,
			LuminosityMin: 3000,
		},
		{
			Name:          "III (Giant)",
			LuminosityMax: 2999,
			LuminosityMin: 0.1,
		},
		{
			Name:          "White Dwarf",
			LuminosityMax: 0.09,
			LuminosityMin: 0.0,
		},
	},
	"B": {
		{
			Name:          "Ia/Ib (Supergiant)",
			LuminosityMax: 999999,
			LuminosityMin: 7500,
		},
		{
			Name:          "II (Bright Giant)",
			LuminosityMax: 7500,
			LuminosityMin: 3000,
		},
		{
			Name:          "III (Giant)",
			LuminosityMax: 2999,
			LuminosityMin: 0.01,
			TempKelvinMax: 12999,
			TempKelvinMin: 10000,
		},
		{
			Name:          "IV (Subgiant)",
			LuminosityMax: 2999,
			LuminosityMin: 0.1,
			TempKelvinMax: 16999,
			TempKelvinMin: 13000,
		},
		{
			Name:          "V (Main Sequence)",
			LuminosityMax: 2999,
			LuminosityMin: 0.1,
			TempKelvinMax: 30000,
			TempKelvinMin: 17000,
		},
		{
			Name:          "White Dwarf",
			LuminosityMax: 0.09,
			LuminosityMin: 0,
		},
	},
	"A": {
		{
			Name:          "Ia/Ib (Supergiant)",
			LuminosityMax: 999999,
			LuminosityMin: 5000,
		},
		{
			Name:          "II (Bright Giant)",
			LuminosityMax: 4999,
			LuminosityMin: 750,
		},
		{
			Name:          "III (Giant)",
			LuminosityMax: 749,
			LuminosityMin: 200,
		},
		{
			Name:          "IV (Subgiant)",
			LuminosityMax: 199,
			LuminosityMin: 50,
		},
		{
			Name:          "V (Main Sequence)",
			LuminosityMax: 49,
			LuminosityMin: 1,
		},
		{
			Name:          "White Dwarf",
			LuminosityMax: 0.9,
			LuminosityMin: 0,
		},
	},
	"F": {
		{
			Name:          "Ia/Ib (Supergiant)",
			LuminosityMax: 999999,
			LuminosityMin: 2000,
		},
		{
			Name:          "II (Bright Giant)",
			LuminosityMax: 1999,
			LuminosityMin: 100,
		},
		{
			Name:          "III (Giant)",
			LuminosityMax: 99,
			LuminosityMin: 40,
		},
		{
			Name:          "IV (Subgiant)",
			LuminosityMax: 39,
			LuminosityMin: 8,
		},
		{
			Name:          "V (Main Sequence)",
			LuminosityMax: 7,
			LuminosityMin: 0.1,
		},
		{
			Name:          "White Dwarf",
			LuminosityMax: 0.09,
			LuminosityMin: 0,
		},
	},
	"G": {
		{
			Name:          "Ia/Ib (Supergiant)",
			LuminosityMax: 999999,
			LuminosityMin: 2500,
		},
		{
			Name:          "II (Bright Giant)",
			LuminosityMax: 2499,
			LuminosityMin: 100,
		},
		{
			Name:          "III (Giant)",
			LuminosityMax: 99,
			LuminosityMin: 15,
		},
		{
			Name:          "IV (Subgiant)",
			LuminosityMax: 14,
			LuminosityMin: 3,
		},
		{
			Name:          "V (Main Sequence)",
			LuminosityMax: 2,
			LuminosityMin: 0.05,
		},
		{
			Name:          "White Dwarf",
			LuminosityMax: 0.04,
			LuminosityMin: 0,
		},
	},
	"K": {
		{
			Name:          "Ia/Ib (Supergiant)",
			LuminosityMax: 999999,
			LuminosityMin: 2500,
		},
		{
			Name:          "II (Bright Giant)",
			LuminosityMax: 2499,
			LuminosityMin: 200,
		},
		{
			Name:          "III (Giant)",
			LuminosityMax: 199,
			LuminosityMin: 15,
		},
		{
			Name:          "IV (Subgiant)",
			LuminosityMax: 14,
			LuminosityMin: 1,
		},
		{
			Name:          "V (Main Sequence)",
			LuminosityMax: 0.9,
			LuminosityMin: 0.001,
		},
		{
			Name:          "White Dwarf",
			LuminosityMax: 0.0009,
			LuminosityMin: 0,
		},
	},
	"M": {
		{
			Name:          "Ia/Ib (Supergiant)",
			LuminosityMax: 999999,
			LuminosityMin: 2500,
		},
		{
			Name:          "II (Bright Giant)",
			LuminosityMax: 2499,
			LuminosityMin: 800,
		},
		{
			Name:          "III (Giant)",
			LuminosityMax: 799,
			LuminosityMin: 1,
		},
		/* {
			Name:          "IV (Subgiant)",
			LuminosityMax: 0.9,
			LuminosityMin: 5,
		}, */
		{
			Name:          "V (Main Sequence)",
			LuminosityMax: 0.9,
			LuminosityMin: 0.00001,
		},
		{
			Name:          "White Dwarf",
			LuminosityMax: 0.000009,
			LuminosityMin: 0,
		},
	},
}
