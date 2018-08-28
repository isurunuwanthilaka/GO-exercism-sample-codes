package space

import (
	"math"
)

func Age(seconds float64, planet string) float64 {

	/*
		Convert seconts to planet years
	*/

	var year float64
	switch planet {
	case "Earth":
		year = (seconds / 31557600)
	case "Mercury":
		year = (seconds / 31557600) / 0.2408467
	case "Venus":
		year = (seconds / 31557600) / 0.61519726
	case "Mars":
		year = (seconds / 31557600) / 1.8808158
	case "Jupiter":
		year = (seconds / 31557600) / 11.862615
	case "Saturn":
		year = (seconds / 31557600) / 29.447498
	case "Uranus":
		year = (seconds / 31557600) / 84.016846
	case "Neptune":
		year = (seconds / 31557600) / 164.79132
	}
	return math.Round(year*100) / 100
}
