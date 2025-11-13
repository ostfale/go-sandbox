package main

import "fmt"

func RunCompositionExample() {
	sep("Composition Example")

	var m = Manager{Employee{"John Denver", "123"}, []Employee{{"Mary Joe", "456"}}}

	fmt.Println(m.ID)
	fmt.Println(m.Description())
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
}
