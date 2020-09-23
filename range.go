package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	// for range like foreach
	// คือ loop ค่าจาก array โดยตรง
	// ยกตัวอย่าง for ตามด้วย ลำดับ หรือ key
	// ตามด้วยค่า value ตามด้วย range และตามด้วยค่า array ที่ต้องการ loop
	// กรณีที่ใส่ _ เพราะไม่ต้องการใช้งานค่าหากต้องการใช้ ให้ใส่เป็นชื่อ
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		// print index เมื่อ num=3
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// for key, velue = range kvs
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	x := 5
	x += 2
}
