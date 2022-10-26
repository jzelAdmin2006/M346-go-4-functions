package main

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	convertCelsiusToFahrenheit(5)   // 41
	convertFahrenheitToCelsius(41)  // 5
	convertCelsiusToFahrenheit(100) // 212
	convertFahrenheitToCelsius(212) // 100
	convertCelsiusToFahrenheit(0)   // 32
	convertFahrenheitToCelsius(32)  // 0
}
