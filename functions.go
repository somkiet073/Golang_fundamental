package main

import "fmt"

// basic function return a+b (integer)
func plus(a int, b int) int {
	return a + b
}

// basic function return a+b+c (integer)
func plusPlus(a, b, c int) int {
	return a + b + c
}

// function Multiple Return Values
func vals() (int, int) {
	return 3, 7
}

// Variadic Functions คือรูปแบบการประกาศ ยฟพฟทห ของ func
// โดยที่ว่าจะส่งตัวแปรเข้ามากี่ตัวก็ได้
// **แต่มีเงื่อนว่า ต้องเป็นตัวแปรประเภทเดียวกันเท่านั้น
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// การประกาศ  anonymous functions
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursive function คือการทำซ้ำ function
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	// Variadic Functions
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//  anonymous functions
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	// basic anonymous functions
	nx := func() int {
		return 12
	}
	fmt.Println(nx())

	// use recursive
	fmt.Println(fact(7))
}
