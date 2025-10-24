package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func RunControlIFStructures() {
	fmt.Println("\nControl Structures overview")

	sep("Control Structures - IF regular")
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("n is zero")
	} else if n >= 5 {
		fmt.Println("n is greater than 5")
	} else {
		fmt.Println("n is less than 5")
	}

	sep("Control Structures - IF with initialization")
	if x := rand.Intn(10); x == 0 {
		fmt.Println("x is zero")
	} else if x >= 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	sep("Control Structures - IF scoped error handling")
	// Initialization statement: 'result, err := checkValue(-5)' ->  Condition: 'err != nil'
	//	if result, err := checkValue(5); err != nil {
	if result, err := checkValue(-5); err != nil {
		// 'err' and 'result' are available here
		fmt.Println("Error:", err)
	} else {
		// 'err' and 'result' are also available here
		fmt.Println("Success:", result)
	}
	// The variables 'result' and 'err' are NOT available here -> fmt.Println(err) This would cause a compile-time error!
}

// A simple function that might fail
func checkValue(value int) (string, error) {
	if value < 0 {
		return "", errors.New("value cannot be negative")
	}
	return "Value is valid", nil
}
