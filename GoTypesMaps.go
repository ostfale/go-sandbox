package main

import (
	"fmt"
	"sort"
)

func GoTypesMaps() {
	sep("Go Types -> Maps Examples")

	// make with a capacity hint
	capitals := make(map[string]string, 4)
	capitals["France"], capitals["Italy"] = "Paris", "Rome"
	capitals["Germany"] = "Berlin"

	// read and presence check
	city := capitals["France"]
	fmt.Println(city)
	missing, ok := capitals["Canada"]
	fmt.Println(ok, missing)

	// delete
	delete(capitals, "Italy")
	fmt.Println(capitals)

	// deterministic iteration: collect and sort keys
	keys := make([]string, 0, len(capitals))
	for k := range capitals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, capitals[k])
	}
}
