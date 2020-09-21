package main

import "fmt"

func main() {

	// สร้าง array มีขนาด 3
	s := make([]string, 3)
	fmt.Println("emp:", s)
	// output emp: [  ]

	// set array 0=a, 1=b, 2=c
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	// output set: [a b c]
	fmt.Println("get:", s[2])
	// output get: c
	fmt.Println("len:", len(s))
	// output len: 3

	// เพิ่ม array d, e, f
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	// output apd: [a b c d e f]

	// สร้าง array โดยมีขนาดเท่ากับ array s
	c := make([]string, len(s))
	// คัดลอก array จาก s มา c
	copy(c, s)
	fmt.Println("cpy:", c)
	// output cpy: [a b c d e f]

	// ตัดตั้งแต่ตำแหน่งที่ 2 ถึงตำแหน่งที่ 4 ไม่เอา 5
	l := s[2:5]
	fmt.Println("sl1:", l)
	// output sl1: [c d e]

	// ตัดตั้งแต่ตำแหน่งที่ 0 ถึง 4 ไม่เอา 5
	l = s[:5]
	fmt.Println("sl2:", l)
	// output sl2: [a b c d e]

	// ประกาศตัวแปร t เป็น array String มีค่าเท่ากับ {g, h, i}
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// output dcl: [g h i]

	// สร้าง array 2 มิติ ขนาด 3
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// กำหนดขนาด array ตามการวนรอบ
		twoD[i] = make([]int, innerLen)
		// วนรอบ append array
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
	// output 2d: [[0] [1 2] [2 3 4]]
}
