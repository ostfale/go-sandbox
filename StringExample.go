package main

import (
	"fmt"
)

func RunStringExample() {
	fmt.Println("String Example")
	sep("String Example - extract single value from string - 1 byte codepoints")
	var s1 = "Hello there"
	var s2 = s1[4:7]
	fmt.Println(s2)

	// longer codepoints
	sep("String Example - length of a string > 1 byte codepoints")
	var scp = "Hello ♥️"
	var length = len(scp)
	var scpSlice = scp[6:]
	fmt.Printf("Code point > 1 byte - length : %d and slice from 4 : %s", length, scpSlice)
}
