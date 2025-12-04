package main

import (
	"fmt"
	"strconv"
)

type UserId int64   // new type definition
type Email = string // alias, exactly the same as string

func GoTypesCustomTypes() {
	sep("Custom Types")
	var id UserId = 42
	fmt.Println("UserID: ", id)

	takesUserId(id)
	takesString(strconv.FormatInt(int64(id), 10)) // explicit conversion to 42

	var e Email = "a@b.org"
	takesString(e)
	fmt.Println(e)
}

func (id UserId) String() string {
	return fmt.Sprintf("user-%d", id)
}

func takesString(s string) {
	fmt.Println(s)
}

func takesUserId(id UserId) {
	fmt.Println(id)
}
