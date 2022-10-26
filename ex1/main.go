package main

func computeGrade(gotPoints, maxPoints float32) float32 {
	return gotPoints/maxPoints*5 + 1
}

func main() {
	computeGrade(0, 6)    // 1
	computeGrade(9.5, 18) // 3.6388888
	computeGrade(10, 10)  // 6
}
