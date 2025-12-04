package main

import "fmt"

type MailCategory int
type State int
type Perms uint8
type Role string

const (
	RoleUnknown Role = "" // 0 - zero value
	RoleAdmin   Role = "admin"
	RoleUser    Role = "user"
)

func GoTypesEnums() {
	sep("Go Types ->  Enums")

	const (
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social
		Updates
	)
	fmt.Println("All enums: ", Uncategorized, Personal, Spam, Social, Updates)

	sep("Go Types ->  Numeric Enums")
	const (
		StateUnknown State = iota // 0 - zero value
		StatePending              // 1
		StateActive               // 2
		StateDone                 // 3
	)
	fmt.Println("Numeric enums: ", StateUnknown, StatePending, StateActive, StateDone)

	sep("Go Types -> Enum Bit Flags")
	const (
		Read    Perms = 1 << iota // 1
		Write                     // 2
		Execute                   // 4
	)

	fmt.Println("Enum bit flags: ", Read, Write, Execute)

	sep("Go Types -> String backend Enums")
	// Role constants moved to package level
	fmt.Println("String backend enums: ", RoleUser, RoleAdmin)
}
func has(p Perms, f Perms) bool { return p&f != 0 }

func (r Role) Valid() bool {
	switch r {
	case RoleUnknown, RoleAdmin, RoleUser:
		return true
	}
	return false
}
