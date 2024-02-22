package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/MrEliasen/sol/src"
	"github.com/MrEliasen/sol/src/config"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	num := int64(rand.Intn(999999999))
	config.SeedRng(num)

	for i := 0; i < 20; i++ {
		system := src.NewStarSystem()
		system.Display()
	}

	fmt.Printf("Seed: %v", num)

	// dist()

	/* earth := &src.Planet{
		System:    sol,
		Radius:    6371,
		Density:   5.5134,
		EarthMass: 1,
	}

	moon := &src.Planet{
		System:    sol,
		Primary:   earth,
		Radius:    1737.4,
		Density:   3.344,
		EarthMass: 7.342e22 / src.EARTH_MASS,
	}

	planets := map[string]*src.Planet{
		"earth": earth,
		"moon":  moon,
	}

	sol.Planets = planets

	// Print the result
	fmt.Printf("Roche Limit for Earth-Moon: %.6f AU\n", sol.Planets["moon"].HillSphere())
	fmt.Printf("Roche Limit for Sun-Earth: %.6f AU\n", sol.Star.RocheLimit(sol.Planets["earth"]))
	fmt.Printf("Hill Sphere Radius for Earth-Moon system: %.6f AU\n", sol.Planets["moon"].HillSphere())
	fmt.Printf("Hill Sphere Radius for Sun-Earth system: %.6f AU\n", sol.Planets["earth"].HillSphere()) */
}

func dist() {
	// Mean and standard deviation of the normal distribution
	mean := 0.0
	stdDev := 1.0

	// Generate random numbers following a normal distribution
	numPoints := 100
	values := make(plotter.Values, numPoints)
	for i := 0; i < numPoints; i++ {
		values[i] = rand.NormFloat64()*stdDev + mean
	}

	// Create a plot
	p := plot.New()

	// Create a histogram of the generated values
	h, err := plotter.NewHist(values, 20)
	if err != nil {
		log.Fatal(err)
	}
	h.Normalize(1) // Normalize the histogram

	// Add the histogram to the plot
	p.Add(h)

	// Save the plot to a file (or display it)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "normal_distribution.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plot saved to normal_distribution.png")
}
