package main

import (
	"oop/employee"
)

func main() {
	e := employee.Employee{
		Firstname:   "Somkiet",
		Lastname:    "Paowang",
		TotalLeaves: 10,
		LeavesToken: 20,
	}
	e.LeavesRemaining()

}
