package main

import (
	"fmt"
	"math"
)

// Sqrt function computes the square root of x using Newton's method
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z - x) / (2*z)
		fmt.Printf("Iteration %d: %f\n", i, z)
	}
	return z
}

// Refined Sqrt function to stop when the value changes by a very small amount
func RefinedSqrt(x float64) float64 {
	z := 1.0
	prevZ := 0.0
	for math.Abs(z - prevZ) > 1e-10 {
		prevZ = z
		z = z - (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println("Approximate square root of 2 (10 iterations):", Sqrt(2))
	fmt.Println("Refined square root of 2:", RefinedSqrt(2))

	// // Testing with other values
	// values := []float64{1, 2, 3, 4, 9, 16, 25}
	// for _, v := range values {
	// 	fmt.Printf("Testing value: %f\n", v)
	// 	fmt.Println("Approximate square root (10 iterations):", Sqrt(v))
	// 	fmt.Println("Refined square root:", RefinedSqrt(v))
	// 	fmt.Println("math.Sqrt result:", math.Sqrt(v))
	// 	fmt.Println()
	// }
}
