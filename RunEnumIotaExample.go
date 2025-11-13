package main

import (
	"fmt"
)

type MailCategory int

func RunEnumIotaExample() {

	sep("Enum Iota Example")
	const (
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social
		Updates
	)

	fmt.Println("All enums: ", Uncategorized, Personal, Spam, Social, Updates)
}
