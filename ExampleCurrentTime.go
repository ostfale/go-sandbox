package main

import (
	"fmt"
	"time"
)

func ExampleCurrentTime() {

	sep("Current Time Example")

	fmt.Println("The current time is: ", time.Now())
}
