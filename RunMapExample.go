package main

import (
	"fmt"
	"maps"
)

func RunMapExample() {
	sep("Map Examples Overview!")

	teams := map[string][]string{
		"England": {"Liverpool", "Manchester", "Chelsea"},
		"Spain":   {"Barcelona", "Real Madrid", "Atletico Madrid"},
		"Germany": {"Bayern Munich", "Bayer Leverkusen", "St. Pauli"},
	}
	fmt.Printf("Print map with international soccer teams\n: %s", teams)

	sep("Map - Accessing Values")
	totalWins := map[string]int{}
	totalWins["England"] = 10
	totalWins["Spain"] = 12
	fmt.Printf("Show England number %d \n", totalWins["England"])
	fmt.Printf("Show Spain number %d \n", totalWins["Spain"])
	totalWins["England"]++
	fmt.Printf("Show England number iterated: %d \n", totalWins["England"])

	sep("Map - Accessing not existing Key using the comma Ok Idiom")
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	v, ok := m["c"]
	fmt.Printf("Value of c is %d and ok is %t \n", v, ok)
	v, ok = m["a"]
	fmt.Printf("Value of a is %d and ok is %t \n", v, ok)

	sep("Delete a key from a map")
	delete(m, "a")
	v, ok = m["a"]
	fmt.Printf("Value of a is %d and ok is %t \n", v, ok)

	sep("Empty Map")
	m = map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("Map length: %d \n", len(m))
	clear(m)
	fmt.Printf("Map length after clear: %d \n", len(m))

	sep("Compare Maps using helper function from maps package")
	mc := map[string]int{
		"hello": 1,
		"world": 2,
	}
	nc := map[string]int{
		"hello": 1,
		"world": 2,
	}
	fmt.Printf("Compare maps: %t \n", maps.Equal(mc, nc))
}
