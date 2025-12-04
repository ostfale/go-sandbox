package main

import (
	"fmt"
	"slices"
)

func RunSlicesExamples() {
	fmt.Println("\nSlices Examples")
	// empty slices
	sep("Empty Slice")
	var xi []int
	fmt.Printf("Is an empty slice nil: %t", xi == nil)

	sep("Compare Slices")
	// compare slices
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c", "d", "e"}

	fmt.Printf("\nCompare equal slices: %t\n", slices.Equal(x, y)) // true
	fmt.Printf("Compare unequal slices: %t\n", slices.Equal(x, z)) // false
	fmt.Println(slices.Equal(x, z))                                // true
	fmt.Printf("Print string slice: %s", s)
	//	fmt.Println(slices.Equal(x,s))        // does not compile

	sep("Append Slice")
	// append function
	sApp := []int{10, 20}
	fmt.Printf("\nInitial slice: %v, Length: %d, Capacity: %d\n", sApp, len(sApp), cap(sApp))
	sApp = append(sApp, 30)
	fmt.Printf("After append:  %v, Length: %d, Capacity: %d\n", sApp, len(sApp), cap(sApp))

	yApp := []int{21, 31, 41}
	sApp = append(sApp, yApp...)
	fmt.Printf("After append another slice:  %v, Length: %d, Capacity: %d\n", sApp, len(sApp), cap(sApp))

	sep("Empty Slice")
	// empty a slice
	sClear := []string{"first", "second"}
	fmt.Println(sClear, len(sClear)) // [first,second] 3
	clear(sClear)
	fmt.Println(sClear, len(sClear)) // [ ] 3

	sep("Make Slice")
	// make examples
	sMake := make([]int, 3)
	sMake1 := make([]int, 3, 5)
	fmt.Println(len(sMake), cap(sMake), len(sMake1), cap(sMake1))

	sep("Slicing Slices")
	xSlice := []string{"a", "b", "c", "d", "e"}
	ySlice := xSlice[:2]
	zSlice := xSlice[1:]
	dSlice := xSlice[1:3]
	eSlice := xSlice[:]
	fmt.Println("xSlice", xSlice)
	fmt.Println("ySlice", ySlice)
	fmt.Println("zSlice", zSlice)
	fmt.Println("dSlice", dSlice)
	fmt.Println("eSlice", eSlice)

	sep("Sharing memory")
	// changing slices by using the same memory
	xSlice[0] = "A"
	fmt.Println("xSlice", xSlice)
	fmt.Println("ySlice", ySlice)

	sep("Copying Slices")
	// copy slices
	xCopy := []int{1, 2, 3, 4, 5}
	yCopy := make([]int, 4)
	num := copy(yCopy, xCopy) // first argument is destination slice, second is source slice
	fmt.Println(yCopy, num)

	// copy slices from an array
	xCopyArray := []int{1, 2, 3, 4}
	dArray := [4]int{5, 6, 7, 8}
	yCopyArray := make([]int, 2)
	copy(yCopyArray, dArray[:])
	fmt.Println(yCopyArray)
	copy(dArray[:], xCopyArray)
	fmt.Println(dArray)

	sep("Converting Array to Slice")
	// convert array to slice
	x1Slice := []int{1, 2, 3, 4}
	x1Array := [4]int(x1Slice)
	smallArray := [2]int(x1Slice)
	x1Slice[0] = 10
	fmt.Println(x1Slice)
	fmt.Println(x1Array)
	fmt.Println(smallArray)

}
