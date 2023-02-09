package conv

import "math"

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/
//func mathRound(value float64) float64 {
//	mathRound := math.Round(value*100) / 100
//}

func FloatRound(value float64) float64 {
	return math.Round(value*100) / 100
}

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller lignende
	// Legger inn formellen Celsius = (Farhrenheit - 32)*(5/9)
	celsius := (value - 32) * (5.0 / 9)
	//	mathRound := math.Round(celsius*100) / 100
	return FloatRound(celsius) // Skal bli 56.67. Kan også brukes i 1 linje

}

func CelsiusToFahrenheit(value float64) float64 { //Farhrenheit = Celsius*(9/5) + 32
	return (value)*(9/5.0) + 32
}

func KelvinToFarhenheit(value float64) float64 { //     Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
	return (value-273.15)*(9/5.0) + 32
}

func FarhenheitToKelvin(value float64) float64 { //(Farhenheit - 32) * (5/9) + 273.15
	return (value-32)*(5.0/9) + 273.15
}

func KelvinToCelsius(value float64) float64 { //Celsius = Kelvin - 273.15
	return value - 273.15
}

func CelsiusToKelvin(value float64) float64 { //Kelvin = Celsius + 273.15
	return value + 273.15
}

// De andre konverteringsfunksjonene implementere her
// ...
