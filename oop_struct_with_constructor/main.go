package main

import (
	"oop_struct_with_constructor/employee"
)

func main() {
	e := employee.Init(
		"Somkiet",
		"Paowang",
		30,
		20,
	)

	e.LeavesRemaining()
}
