package main

import (
	"flag"
	"fmt"
	"github.com/SiMENhol/funtemps/conv"
	"strings"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfact string
var temp float64

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.Float64Var(&temp, "Temperatur", 0.0, "Temperatur ute")

}

func main() {
	/*
		fmt.Println(conv.KelvinToCelsius(20))
		fmt.Println(funfacts.GetFunFacts("sun"))
		fmt.Println(funfacts.GetFunFacts("luna"))
		fmt.Println(funfacts.GetFunFacts("terra"))
	*/
	flag.Parse()

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, cels, kelv, out, funfact, temp) //Printer ut forskjellige verdier

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Konvertere temperaturene
	if out == "C" && isFlagPassed("F") {
		fmt.Printf("%.f Fahrenheit er %s Celsius\n", fahr, trimDecimal(conv.FarhenheitToCelsius(fahr))) //, addSpaceSeperator(strconv.FormatFloat(fahr, 2, 64)))
	} else if out == "K" && isFlagPassed("F") {
		fmt.Printf("%.f Fahrenheit er %s Kelvin \n", fahr, trimDecimal(conv.FarhenheitToKelvin(fahr)))
	} else if out == "F" && isFlagPassed("C") {
		fmt.Printf("%.f Celsius er %s Fahrenheit \n", cels, trimDecimal(conv.CelsiusToFahrenheit(cels)))
	} else if out == "K" && isFlagPassed("C") {
		fmt.Printf("%.f Celsius er %s Kelvin \n", cels, trimDecimal(conv.CelsiusToKelvin(cels)))
	} else if out == "C" && isFlagPassed("K") {
		fmt.Printf("%.f Kelvin er %s Celsius \n", kelv, trimDecimal(conv.KelvinToCelsius(kelv)))
	} else if out == "F" && isFlagPassed("K") {
		fmt.Printf("%.f Kelvin er %s Fahrenheit \n", kelv, trimDecimal(conv.KelvinToFarhenheit(kelv)))
	}

}

// Funksjone for å fjerne desimal hvis det er .00
func trimDecimal(value float64) string {
	str := fmt.Sprintf("%.2f", value)
	return strings.TrimRight(strings.TrimRight(str, "0"), ".")
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
