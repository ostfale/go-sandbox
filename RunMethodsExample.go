package main

import (
	"fmt"
	"time"
)

func RunMethodsExample() {
	fmt.Println("\nMethods Examples (Pointer and Value Receivers)")
	sep("Methods Examples - Pointer Receiver")
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	// Total: 0, Last Updated: 0001-01-01 00:00:00 +0000 UTC
	// Total: 1, Last Updated: 2025-11-11 07:49:52.8352778 +0100 CET m=+0.135077501

	sep("Methods Examples - nil Receiver")

	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(11)) // false
}

// nil handlers

type IntTree struct {
	Value       int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val, nil, nil}
	}

	if val < it.Value {
		it.left = it.left.Insert(val)
	}
	if val > it.Value {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.Value:
		return it.left.Contains(val)
	case val > it.Value:
		return it.right.Contains(val)
	default:
		return true
	}
}

// pointer and value receivers

type Counter struct {
	total       int
	lastUpdated time.Time // The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("Total: %d, Last Updated: %v", c.total, c.lastUpdated)
}
