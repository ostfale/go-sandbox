package main

import (
	"fmt"
)

func RunPointerExamples() {
	fmt.Println("\n\nPointer Examples")

	sep("Pointer Address Examples")
	pointerAddress()
}

func pointerAddress() {
	i := 10 // A regular integer variable
	p := &i // p is a pointer to i (stores i's memory address)

	fmt.Println("Value of i:", i)             // Output: 10
	fmt.Println("Address of i:", p)           // Output: (a memory address, e.g., 0xc000014080)
	fmt.Println("Value pointed to by p:", *p) // Output: 10 (Dereferencing p)

	*p = 20 // Change the value at the memory address p points to

	fmt.Println("New value of variable is:", i) // Output: 20 (i's value has changed)

	// pointer type
	xpType := 10
	var pointerToXPType *int
	pointerToXPType = &xpType
	fmt.Println("\nPointer to xpType:", pointerToXPType)
	fmt.Println("Value pointed to by pointerToXPType:", *pointerToXPType)
	*pointerToXPType = 20
	fmt.Println("Value pointed to by pointerToXPType:", *pointerToXPType)
	fmt.Println("Value of xpType:", xpType)
}
