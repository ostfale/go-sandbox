package main

import "fmt"

func GoTypeConstants() {
	sep("Go Type Constants")

	const PI = 3.14159265 // float64
	const Answer = 42     // untyped integer constant
	const Eacute = 'é'    // untyped rune constant

	var r int64 = Answer
	var f float32 = Answer
	var ch rune = Eacute
	fmt.Printf("Constants: %T %T %T %v\n", r, f, ch, PI) // Constants: int64 float32 int32 3.14159265

}
