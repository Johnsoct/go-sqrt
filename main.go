package main

import (
	"fmt"
	"math"
)

func isAcceptableDeviation(a, b float64) bool {
	deviance := 2.220446049250313e-16
	return math.Abs(a-b) > deviance
}

func newtonsApproximation(z, x float64) float64 {
	return ((z*z - x) / (2 * z))
}

func Sqrt(x float64) float64 {
	iterations := 0
	last := float64(0)
	next := float64(1)

	for isAcceptableDeviation(next, last) {
		last = next
		next -= newtonsApproximation(next, x)
		iterations++
	}

	// fmt.Printf("Iterations:\t%d\n", iterations)
	return last
}

func main() {
	for i := 1.0; i < 4.0; i++ {
		fmt.Println("Sqrt", i, " : ", Sqrt(i))
		fmt.Println("Math", i, " : ", math.Sqrt(i))
		fmt.Println("\n")
	}
}
