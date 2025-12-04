package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func RunStringExample() {
	fmt.Println("\nString Example")
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

	StringTypeExample()
	BytesTypeExample()
	RunesTypeExample()
}

func StringTypeExample() {
	sep("String Type Example")

	s := "Hello, cafe"  // UTF-8 encoded string
	fmt.Println(len(s)) // 12 bytes

	sub := s[7:]
	fmt.Println(sub) // "cafe"

	copied := "" + sub // force a copy
	fmt.Println(copied)
}

func BytesTypeExample() {
	sep("Bytes Type Example")

	s := "go"
	b := []byte(s) // copy
	b[0] = 'G'     // mutate

	var buff bytes.Buffer
	buff.Write(b)
	buff.WriteString(" lang")
	out := buff.String()
	fmt.Println(out)
}

func RunesTypeExample() {
	sep("Runes Type Example")

	s := "Hello, café"
	fmt.Println(len(s), utf8.RuneCountInString(s)) // 12 bytes, 11 runes

	for i, r := range s { // safe iteration by runes
		fmt.Printf("%d:%c", i, r)
	}
	fmt.Println()

	rs := []rune(s) // edit by characters
	rs[0] = 'h'
	fmt.Println(string(rs)) // hello, café
}
