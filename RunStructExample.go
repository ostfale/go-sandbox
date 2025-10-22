package main

import (
	"fmt"
)

func RunStructExample() {
	fmt.Println("\nStruct Examples")

	sep("Struct Initialization")
	type Person struct {
		Name    string
		Age     int
		IsAdult bool
	}

	// initialize
	p1 := Person{"Alice", 30, true} // init class
	fmt.Println(p1)
	p2 := Person{Name: "Bob", Age: 25} // using field names

	// accessing
	fmt.Println(p2) // Output: Bob
	p2.Age = 26     // Modifying a field
	fmt.Printf("Bobs age incremented: %d", p2.Age)

	sep("Anonymous Structs")
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "Louis"
	person.age = 15
	person.pet = "dog"
	fmt.Println(person)
	pet := struct {
		name string
		kind string
	}{
		name: "Cerberus",
		kind: "dog",
	}
	fmt.Println(person, pet)
}
