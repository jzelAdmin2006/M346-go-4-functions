package main

import (
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
	a float64
	b float64
}

func (shortSides ShortSides) getHypotenuse() float64 {
	return math.Sqrt(math.Pow(shortSides.a, 2) + math.Pow(shortSides.b, 2))
}

func main() {
	computeHypotenuse(3, 4)               // 5
	computeHypotenuse(1, 1)               // 1.4142135623730951
	computeHypotenuse(3264.467, 31278.23) // 31448.122626430166

	var firstShortSides ShortSides = ShortSides{3, 4}
	firstShortSides.getHypotenuse() // 5

	var secondShortSides ShortSides = ShortSides{1, 1}
	secondShortSides.getHypotenuse() // 1.4142135623730951

	var thirdShortSides ShortSides = ShortSides{3264.467, 31278.23}
	thirdShortSides.getHypotenuse() // 31448.122626430166
}
