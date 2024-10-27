package main

import (
	"fmt"
	"math"
)

func isAcceptableDeviation(guess, lastGuess float64) bool {
	deviance := 2.220446049250313e-16
	return math.Abs(guess - lastGuess) <= deviance
}

func newtonsMethod(z, x float64) float64 {
	newGuess := z - ((z*z - x) / (2 * z))
	return newGuess
}

func Sqrt(x float64) float64 {
	count := 0
	guess := float64(1)
	previousGuess := float64(0)

	for i := 1; !isAcceptableDeviation(guess, previousGuess); i++ {
		previousGuess = guess
		guess = newtonsMethod(guess, x)
		count++
	}

	// fmt.Printf("Count:\t%d\n", count)
	return previousGuess
}

func main() {
	for i := 1.0; i < 4.0; i++ {
		fmt.Println("Sqrt of ", i, " is ", Sqrt(i))
		fmt.Println("Math Sqrt of ", i, " is ", math.Sqrt(i))
	}
}
