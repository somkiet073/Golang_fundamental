package main

import "fmt"

/*  ประเภทของตัวแปร Go มีค่าประเภทต่างๆ เช่น
String, Integers, floats, booleans, etc.
ประเภทเหล่านี้คือ ประเภทพื้นฐานที่ใช้งานบ่อย
*/
func main() {

	// String
	fmt.Println("go" + "lang")

	// Integers
	fmt.Println("1+1 =", 1+1)
	// floates
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}
