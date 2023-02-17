package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
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
	/*
	   Her er eksempler på hvordan 	man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.Float64Var(&temp, "Temperatur", 0.0, "Temperatur ute")

}

func main() {
	//fmt.Println(conv.KelvinToCelsius(20))
	//fmt.Println(funfacts.GetFunFacts("sun"))
	//fmt.Println(funfacts.GetFunFacts("luna"))
	//fmt.Println(funfacts.GetFunFacts("terra"))
	flag.Parse()

	/*
		    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
		    pakkene implementeres.

		    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
		    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
		    hvor mange flagg og argumenter er spesifisert på kommandolinje.

		        fmt.Println("len(flag.Args())", len(flag.Args()))
				    fmt.Println("flag.NFlag()", flag.NFlag())

		    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
		    brukes for å utelukke ugyldige kombinasjoner:
		    -F, -C, -K kan ikke brukes samtidig
		    disse tre kan brukes med -out, men ikke med -funfacts
		    -funfacts kan brukes kun med -t

	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, cels, kelv, out, funfact, temp)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Printf("%gF er %gC", fahr, conv.FarhenheitToCelsius(fahr)) //, addSpaceSeperator(strconv.FormatFloat(fahr, 2, 64)))
	} else if out == "K" && isFlagPassed("F") {
		fmt.Printf("%gF er %gK", fahr, conv.FarhenheitToKelvin(fahr))
	} else if out == "F" && isFlagPassed("C") {
		fmt.Printf("%gC er %gF", cels, conv.CelsiusToFahrenheit(cels))
	} else if out == "K" && isFlagPassed("C") {
		fmt.Printf("%gC er %gK", cels, conv.CelsiusToKelvin(cels))
	} else if out == "C" && isFlagPassed("K") {
		fmt.Printf("%gK er %gC", kelv, conv.KelvinToCelsius(kelv))
	} else if out == "F" && isFlagPassed("K") {
		fmt.Printf("%gK er %gF", kelv, conv.KelvinToFarhenheit(kelv))
	}
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
