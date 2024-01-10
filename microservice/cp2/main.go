package main

import (
	"fmt"
	"os"
	"strconv"
)

func celsius2fahrenhiet(t float64) float64 {
	return 9.0/5.0*t + 32
}

func fahrenheit2celsius(t float64) float64 {
	return (t - 32) * 5.0 / 9.0
}

func usage() {
	fmt.Println("Usage: cp2 <temp> <unit>")
	fmt.Println("where <unit> is one of C, F")
	fmt.Println("'mode' is optional, and defaults to 'C', 'F' is also supported")
	fmt.Println()
	fmt.Println("'temperature' is a floating point number to be converted according to mode")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	mode := os.Args[1]
	if mode != "C" && mode != "F" {
		usage()
	}

	float, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		usage()
	}
	var converted float64
	if mode == "C" {
		converted = celsius2fahrenhiet(float)
	} else {
		converted = fahrenheit2celsius(float)
	}
	fmt.Println(converted)
}
