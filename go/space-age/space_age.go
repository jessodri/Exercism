package space

import (
	"fmt"
)

// Planet struct
type Planet string

// Age function calculates how old you are on another planet when given an age in seconds
func Age(seconds float64, planet Planet) (age float64) {

	earthYear := seconds / 31557600

	switch {
	case planet == "Earth":
		age = earthYear
	case planet == "Mercury":
		age = earthYear / 0.2408467
	case planet == "Venus":
		age = earthYear / 0.61519726
	case planet == "Mars":
		age = earthYear / 1.8808158
	case planet == "Jupiter":
		age = earthYear / 11.862615
	case planet == "Saturn":
		age = earthYear / 29.447498
	case planet == "Uranus":
		age = earthYear / 84.016846
	case planet == "Neptune":
		age = earthYear / 164.79132
	}

	fmt.Printf("Your age on %s is %.2f years\n", planet, age)
	return age
}
