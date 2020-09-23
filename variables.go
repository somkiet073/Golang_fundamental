package main

import (
	"fmt"
	"time"
)

/*
	Golang Data Types
*/

func main() {
	/*
		ประกาศตัวแปร a Data Types String
		ตัวแปร aa Data Types string ที่มีการประกาศแล้วไม่มีค่าตั้งต้น
		ค่าที่ได้จะเป็น null
	*/
	var a = "initial"
	fmt.Println(a)
	var aa string
	fmt.Println(aa)

	// ประกาศตัวแปร b, c Data Types Integer
	var b, c int = 1, 2
	fmt.Println(b, c)
	// ประกาศตัวแปร e Data Types Integer ไม่มีการใส่ค่าให้ ค่าที่ได้จะเป็น 0
	var d int
	fmt.Println(d)

	// ประกาศตัวแปร d Data Types boolean
	var e = true
	fmt.Println(e)

	/*
		ประกาศตัวแปร f Data Types String
		:= คือการประกาศตัวแปรแบบ รูปย่อ เมื่อเขียนแบบเต็มจะได้ ดังนี้
		var f string = "apple"
	*/
	f := "apple"
	fmt.Println(f)

	times := time.Second * 2
	fmt.Println("times:", times)
}
