package main

import "fmt"

func main() {
	i := 42 						// int
	fmt.Printf("i is of type %T\n", i)

	f := 3.142					// float
	fmt.Printf("f is of type %T\n", f)

	g := 0.867 + 0.5i		// complex128
	fmt.Printf("g is of type %T\n", g)
}