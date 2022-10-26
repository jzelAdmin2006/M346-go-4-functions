package main

import (
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	computeHypotenuse(3, 4)               // 5
	computeHypotenuse(1, 1)               // 1.4142135623730951
	computeHypotenuse(3264.467, 31278.23) // 31448.122626430166
}
