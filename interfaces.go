package main

import (
	"fmt"
	"math"
)

// interface คือ การระบุว่า struct นั้นๆ ต้องมี func อะไรบ้าง
// คล้าย กับ implement class ของภาษา PHP
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// add func area() เข้าไปใน rect struct
func (r rect) area() float64 {
	return r.width * r.height
}

// add func perim() เข้าไปใน rect struct
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// add func area() เข้าไปใน circle struct
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// add func perim() เข้าไปใน circle struct
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// func call geometry interface
func mesure(g geometry) {
	// print ดูค่า ของ Struct
	fmt.Println(g)
	// print ดูค่า ของ func area ใน struct นั้นๆ
	fmt.Println(g.area())
	// print ดูค่า ของ func perim ใน struct นั้นๆ
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	mesure(r)
	mesure(c)
}
