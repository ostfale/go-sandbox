package main

import (
	"fmt"
	"strconv"
)

type opFuncType func(int, int) int

func RunCalculatorExample() {
	fmt.Println("\n\nCalculator Example")

	var opMap = map[string]opFuncType{
		"+": addi,
		"-": sub,
		"*": multi,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Invalid operator", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
		}

		if p2 == 0 && op == "/" {
			fmt.Println("Division by zero")
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(p1, op, p2, "=", result)
	}
}

func addi(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func multi(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}
