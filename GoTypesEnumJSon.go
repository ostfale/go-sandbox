package main

import (
	"encoding/json"
	"fmt"
)

const (
	StateUnknown = iota
	StatePending
	StateActive
	StateDone
)

func GoTypesEnumJSon() {
	sep("Go Types ->  Enums -> JSON")

	var s State = StateActive
	marshal, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(marshal))

}

func (s State) String() string {
	switch s {
	case StateUnknown:
		return "Unknown"
	case StatePending:
		return "Pending"
	case StateActive:
		return "Active"
	case StateDone:
		return "Done"
	default:
		return "State (" + fmt.Sprint(int(s)) + ")"
	}

}

func (s State) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String)
}

func (s *State) UnmarshalJSON(b []byte) error {
	var name string

	if err := json.Unmarshal(b, &name); err != nil {
		return err
	}

	switch name {
	case "Unknown":
		*s = StateUnknown
	case "Pending":
		*s = StatePending
	case "Active":
		*s = StateActive
	case "Done":
		*s = StateDone
	default:
		return fmt.Errorf("invalid State: %s", name)
	}
	return nil
}
