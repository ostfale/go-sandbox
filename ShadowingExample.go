package main

import "fmt"

func RunShadowingExample() {
	fmt.Println("\nShadowing Example")

	sep("Basic Shadowing Example")
	x := 10 // Outer scope variable 'x'
	if true {
		fmt.Println("Inner before shadowing x:", x) // Output: Inner x: 20
		x := 20                                     // Inner scope, new variable 'x' is declared.
		// This 'x' shadows the outer 'x'.
		fmt.Println("Inner after shadowing x:", x) // Output: Inner x: 20
	}
	fmt.Println("Outer x:", x) // Output: Outer x: 10
}
