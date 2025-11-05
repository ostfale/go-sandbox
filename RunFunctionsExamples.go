package main

import (
	"errors"
	"fmt"
	"sort"
)

func RunFunctionsExamples() {
	fmt.Println("\n +++++++++++++++ Function Examples +++++++++++++++++++++")

	sep("Variadic Function Examples")
	sum(1, 2)
	sum()
	sum(5, 10, 15, 20)

	sep("Function with multiple return values")
	result, success := operation(10, 2)
	fmt.Println("Result:", result, "Success:", success) // Output: Result: 5 Success: true

	// Common Go practice: checking the boolean 'ok' value
	if _, ok := operation(10, 0); !ok {
		fmt.Println("Error: Division by zero!")
	}

	sep("Function with named return values")
	result, remainder, err := operationWithNamedParameters(10, 2)
	fmt.Println("Result:", result, "Remainder:", remainder, "Error:", err)

	sep("Function are values")
	useFunctionAsVariable()

	sep("Function type declaration")
	type MathOperation func(int, int) int
	var op MathOperation
	op = add
	fmt.Println("Addition:", op(5, 3)) // Output: Addition: 8
	op = multiply
	fmt.Println("Multiplication:", op(5, 3)) // Output: Multiplication: 15

	sep("Anonymous Functions")
	anonymousFunction()

	sep("Closures - Function Factory with stateful counter")
	counter1 := createCounter()
	fmt.Println("Increment Counter 1:", counter1()) // Output: 1 (count is now 1)
	fmt.Println("Increment Counter 1:", counter1()) // Output: 2 (count is now 2)
	// 2. Calling createCounter *again* creates a *new*, independent 'count' variable (initially 0)
	counter2 := createCounter()
	fmt.Println("Increment Counter 2:", counter2()) // Output: 1 (It has its own separate count)
	fmt.Println("Increment Counter 1:", counter1()) // Output: 3 (The first counter's state is preserved and continues)

	sep("Functions as parameters")
	functionAsParameter()

	sep("Return a closure from a function")
	returnClosureFromFunction()

	sep("Functions - defer example -> LIFO")
	deferExample()
}

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("Defer first value:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("Defer second value:", val)
	}(a)
	a = 30
	defer func(val int) {
		fmt.Println("Defer third value:", val)
	}(a)
	return a
}

func returnClosureFromFunction() {
	twoBase := makeMultiplier(2)
	threeBase := makeMultiplier(3)
	for i := 0; i < 5; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func functionAsParameter() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Jane", "Doe", 62},
		{"John", "Toe", 30},
	}

	fmt.Println("People:", people)

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("Sorted: People:", people)

	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Age Sorted: People:", people)
}

func createCounter() func() int {
	count := 0 // This variable is defined in the outer function's scope

	// The inner, anonymous function is the closure
	return func() int {
		count++ // It accesses and modifies 'count'
		return count
	}
}

func operation(a, b int) (int, bool) {
	if b == 0 {
		// Return a zero value and false if the operation fails
		return 0, false
	}
	// Return the result and true if the operation succeeds
	return a / b, true
}

func anonymousFunction() {
	f := func(j int) {
		fmt.Println("printing:", j, " from outside loop")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}

func useFunctionAsVariable() {
	var mathOperation func(int, int) int // Declare a variable that can hold a function that takes two 'int's and returns an 'int'.
	mathOperation = add                  // Assign the 'add' function to the variable
	result := mathOperation(10, 5)       // Call the function and store the result in 'result'
	fmt.Println("10 + 5 =", result)
}

func add(a, b int) int { // This is the function
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func operationWithNamedParameters(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("division by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err

}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)
}
