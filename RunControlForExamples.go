package main

import (
	"fmt"
	"time"
)

func RunControlForExamples() {
	fmt.Println("\nControl Structures -  For Examples")

	sep("Loop :: C-Style ")
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	sep("Loop :: Condition Style")
	i := 1
	for i <= 100 {
		fmt.Printf("  %d ", i)
		i = i * 2
	}

	sep("Loop :: Continue and Break")
	fmt.Println("Numbers from 1 to 100, skipping 5: and finishing 11")

	for i := 1; i <= 100; i++ {
		// Check if the current number is 5
		if i == 5 {
			// The 'continue' statement immediately jumps to the next iteration
			// of the loop, skipping any code that follows it (in this case, fmt.Println(i)).
			fmt.Println("Skipping 5")
			continue
		}

		if i == 11 {
			fmt.Println("Finishing 11")
			break
		}
		fmt.Println(i)
	}

	sep("Loop :: For Range - Slices")
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		fmt.Printf(" index: %d - evenVals: %d ", i, v)
	}
	fmt.Println()

	sep("Loop :: For Range - Maps")
	capitals := map[string]string{"France": "Paris", "Japan": "Tokyo", "USA": "Washington", "Germany": "Berlin"}

	for country, city := range capitals {
		fmt.Println(country, "->", city)
	}

	sep("Loop :: For Range - Strings")
	for i, v := range "Hello World" {
		fmt.Printf("index: %d  - %c ", i, v)
	}
	fmt.Println()

	sep("Switch statement")
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("%s is a short word\n", word)
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
			fmt.Printf("%s is a long word\n", word)
		default:
			fmt.Printf("%s is a very long word\n", word)
		}
	}

	sep("Switch - Special type blank switch")
	t := time.Now()
	switch { // No expression here; defaults to 'switch true'
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() > 17:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good afternoon")
	}

}
