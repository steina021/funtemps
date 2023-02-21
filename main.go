package main

import (
	"flag"
	"fmt"
	"main/Conv" // TODO: importere funksjonene fra Conv og funfacts
	"math"
	"os"
	"strings"
)

// Definerer flag-variablene i hoved-"scope"
// TODO:
var fahr float64
var cels float64
var kelv float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.00, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.00, "temperatur i grader Celsius")
	flag.Float64Var(&kelv, "K", 0.00, "temperatur i grader Kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

}

func main() {

	flag.Parse()

	/**
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
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	//fmt.Println(fahr, out, funfacts)
	//fmt.Println(cels, out)
	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())
	//fmt.Println(isFlagPassed("out"))
  var desi int
  arguments := strings.Split(os.Args[2], ".")
  fmt.Println(arguments)
  if len(arguments) <= 1 {
  fmt.Println("") 
  } else {
    desi = (len(arguments[1]))
  }
	// fmt.Println(strings.Split(lenght, "."))

	if flag.NFlag() < 2 {
		fmt.Println("expected two flags eks. funtemps -C 100 -out F")
		os.Exit(1)
	}
	if isFlagPassed("out") == false {
		fmt.Println("You must include an out")
	}

	// logikk
	if out == "K" && isFlagPassed("C") {
		// celsius til kelvin
    result := math.Pow10(desi)
		answer := conv.CelsiusToKelvin(cels)
    answer = (math.Round(answer*float64(result))/float64(result))
	  test := fmt.Sprintf("%[1]g°C er %[2]g°K\n", cels, answer)
    fmt.Println(test)
	} else if out == "K" && isFlagPassed("F") {
		// fahrenheit to kelvin
    result := math.Pow10(desi)
		answer := conv.FahrenheitToKelvin(fahr)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°F er %[2]g°K\n", fahr, answer)
    fmt.Println(test)
	} else if out == "F" && isFlagPassed("C") {
		// celsius til fahrenheit
     result := math.Pow10(desi)
		answer := conv.CelsiusToFahrenheit(cels)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°C er %[2]g°F\n", cels, answer)
    fmt.Println(test)
	} else if out == "F" && isFlagPassed("K") {
		// kelvin til fahrenheit
     result := math.Pow10(desi)
		answer := conv.KelvinToFahrenheit(kelv)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°K er %[2]g°F\n", kelv, answer)
    fmt.Println(test)
	} else if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
    result := math.Pow10(desi)
		answer := conv.FahrenheitToCelsius(fahr)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°F er %[2]g°C\n", fahr, answer)
    fmt.Println(test)
	} else if out == "C" && isFlagPassed("K") {
		//Kelvin til celsius
    result := math.Pow10(desi)
		answer := conv.KelvinToCelsius(kelv)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°K er %[2]g°C\n", kelv, answer)
    fmt.Println(test)
	} else {
    fmt.Println("Usage: ")
  }
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
