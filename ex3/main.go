package main

import (
	"math"
)

func computeQuadraticFormula(a, b, c float64) []float64 {
	var discriminant float64 = computeDiscriminant(a, b, c)
	if discriminant > 0 {
		var discriminantSqrt float64 = math.Sqrt(discriminant)
		return []float64{(-b + discriminantSqrt) / a / 2, (-b - discriminantSqrt) / a / 2}
	} else if discriminant == 0 {
		return []float64{-b / a / 2}
	} else {
		return []float64{}
	}
}

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func main() {
	computeQuadraticFormula(3, 4, 1) // [-0.3333333333333333 -1]
	computeQuadraticFormula(2, 4, 2) // [-1]
	computeQuadraticFormula(3, 4, 2) // []

	computeDiscriminant(3, 4, 1) // 4
	computeDiscriminant(2, 4, 2) // 0
	computeDiscriminant(3, 4, 2) // -8
}
