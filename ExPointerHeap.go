package main

import (
	"fmt"
)

func ExPointerHeap() {
	fmt.Println("\nExercise Pointer ::  Heap examples")
	MakePerson("John", "Doe", 30)
	MakePersonPointer("Jane", "Doe", 29)
}

func MakePerson(firstName, lastName string, age int) Person {
	fmt.Println("MakePerson from arguments and return Person object")
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	fmt.Println("MakePersonPointer from arguments and return Person object")
	return &Person{firstName, lastName, age}
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
