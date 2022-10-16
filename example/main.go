package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func f(x int) int {
	return 2 * x
}

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(whom string) {
	fmt.Println("Hello,", whom)
}

func outputCurrency(amount float32, currency rune) {
	fmt.Printf("%8.2f %c\n", amount, currency)
}

func formatCurrency(amount float32, currency rune) string {
	return fmt.Sprintf("%8.2f %c", amount, currency)
}

func rollDice() int {
	return rand.Intn(6) + 1
}

func divide(dividend, divisor float32) (float32, error) {
	if divisor == 0.0 {
		return 0.0, errors.New("divide by 0")
	}
	return dividend / divisor, nil
}

func computeAverage(values []float32) (float32, error) {
	if len(values) == 0 {
		return 0.0, fmt.Errorf("cannot compute average of %v", values)
	}
	var sum float32
	for _, value := range values {
		sum += float32(value)
	}
	return sum / float32(len(values)), nil
}

func makeRandomGrades() []float32 {
	grades := make([]float32, 0)
	nGrades := rand.Intn(3) // 0, 1, or 2
	for i := 0; i < nGrades; i++ {
		grades = append(grades, float32(rand.Intn(500)+100)/100.0)
	}
	return grades
}

type Celsius float32

func outputCelsius(c Celsius) {
	fmt.Printf("%.2f°C\n", c)
}

func (c Celsius) Output() {
	fmt.Printf("%.2f°C\n", c)
}

func main() {
	xs := []int{-2, -1, 0, 1, 2}
	for _, x := range xs {
		y := f(x) // the actual function call
		fmt.Printf("y=f(%d)=%d\n", x, y)
	}

	sayHello()

	sayHelloTo("Alice")

	outputCurrency(2.5, '$')
	outputCurrency(10.0/3.0, '€')
	outputCurrency(1234.567, '¥')

	dollars := formatCurrency(2.5, '$')
	euros := formatCurrency(10.0/3.0, '€')
	fmt.Println(dollars)
	fmt.Println(euros)

	rand.Seed(time.Now().Unix())
	fmt.Println(rollDice())
	fmt.Println(rollDice())
	fmt.Println(rollDice())

	fmt.Println(divide(10.0, 3.0))
	fmt.Println(divide(10.0, 0.0))

	grades := makeRandomGrades()
	average, err := computeAverage(grades)
	if err != nil {
		fmt.Fprintf(os.Stderr, "compute average of %v: %v\n", grades, err)
	} else {
		fmt.Printf("the average of %v is %.2f\n", grades, average)
	}

	var coldest Celsius = -273.15
	var warm Celsius = 32.5

	outputCelsius(coldest)
	outputCelsius(warm)
	coldest.Output()
	warm.Output()
}
