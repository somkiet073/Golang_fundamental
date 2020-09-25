package main

import (
	"example2/auth"
	"fmt"
)

func main() {
	username := "admin"
	password := "1234"
	tokenKey := "12345678"

	chk := auth.CheckLogin(
		username,
		password,
		tokenKey,
	)
	xx, yy := fmt.Scanf("Enter word:")
	fmt.Println(chk)
	fmt.Println(xx)
	fmt.Println(yy)
}
