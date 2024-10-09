package main

import (
	"fmt"
)

type MyError struct {
	What string
	Input float64
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %s %v", e.What, e.Input)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError{What: "cannot Sqrt negative number: ", Input: x}
	} else {
		z := 1.0
		for i := 0; i < 10; i++ {
			z = z - (z*z-x)/(2*z)
		}
		return z, nil
	}
}

func run(x float64) error {
	if x < 0 {
		return &MyError{
			What: "cannot Sqrt negative number: ",
			Input: x,
		}
	}
	return nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}