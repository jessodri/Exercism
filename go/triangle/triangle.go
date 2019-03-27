// Package triangle determines the type of triange
package triangle

import (
	"log"
	"math"
)

// Kind is the type of triange
type Kind int

const (
	// NaT not a triangle
	NaT = iota
	// Equ equilateral - All 3 sides are equal
	Equ
	// Iso isosceles - Two sides are equal
	Iso
	// Sca scalene - All sides are different lengths
	Sca
)

// KindFromSides is a function determining if 3 lengths will make a triangle
func KindFromSides(a, b, c float64) Kind {

	equilateral := a == b && a == c && b == c
	isosceles := a == b || b == c || c == a
	scalene := a != b && a != c && b != c
	degenerate := a+b == c || a+c == b || c+b == a

	notANumber := math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
	notFinite := math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
	sidesEqualZero := a <= 0 || b <= 0 || c <= 0
	sumTwoSidesLessThanThird := a+b < c || a+c < b || c+b < a

	notTriangle := notANumber || notFinite || sidesEqualZero || sumTwoSidesLessThanThird

	var k Kind

	if notTriangle {
		log.Println("Not a Triangle")
		return k
	}

	if scalene {
		log.Println("Scalene")

		k = Sca
	}

	if equilateral {
		log.Println("Equilateral")

		k = Equ
	}

	if isosceles && !equilateral {
		log.Println("Isosceles")

		k = Iso
	}

	if degenerate {
		log.Println("degenerate")
	}

	return k
}
