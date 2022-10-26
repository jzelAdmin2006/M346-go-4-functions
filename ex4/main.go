package main

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

type Celsius float64

func (celsius Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(float64(celsius)*9/5 + 32)
}

type Fahrenheit float64

func (fahrenheit Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((float64(fahrenheit) - 32) * 5 / 9)
}

func main() {

	convertCelsiusToFahrenheit(5)   // 41
	convertFahrenheitToCelsius(41)  // 5
	convertCelsiusToFahrenheit(100) // 212
	convertFahrenheitToCelsius(212) // 100
	convertCelsiusToFahrenheit(0)   // 32
	convertFahrenheitToCelsius(32)  // 0

	var cozy Celsius = 23.0
	cozy.ConvertToFahrenheit() // 73.4

	var cold Fahrenheit = 15.3
	cold.ConvertToCelsius() // -9.277777777777779

	var quiteColdCelsius Celsius = 5
	var waterBoilingTemperatureCelsius Celsius = 100
	var waterSolidificationTemperatureCelsius Celsius = 0

	var quiteColdFahrenheit Fahrenheit = quiteColdCelsius.ConvertToFahrenheit()                                           // 41
	quiteColdFahrenheit.ConvertToCelsius()                                                                                // 5
	var waterBoilingTemperatureFahrenheit Fahrenheit = waterBoilingTemperatureCelsius.ConvertToFahrenheit()               // 212
	waterBoilingTemperatureFahrenheit.ConvertToCelsius()                                                                  // 100
	var waterSolidificationTemperatureFahrenheit Fahrenheit = waterSolidificationTemperatureCelsius.ConvertToFahrenheit() // 32
	waterSolidificationTemperatureFahrenheit.ConvertToCelsius()                                                           // 0
}
