package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func changeName(p person) {
	p.name = "Bob"
}

func changeNamePPointer(p *person) {
	p.name = "Bob"
}

func main() {
	fmt.Println("Person 1:", person{"Bob", 20})
	fmt.Println("Person 2:", person{name: "Alice", age: 30})
	fmt.Println("Person 3:", person{name: "fRED"})
	fmt.Println("Person 4:", &person{name: "Ann", age: 40})

	fmt.Println("Person 5:", newPerson("jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println("person 6:", s.name)

	sp := &s
	fmt.Println("person 7:", sp.age)

	sp.age = 51
	fmt.Println("person 7:", sp.age)

	/*
		Go pass by value ทุกๆครั้งที่เราส่งพารามิเตอร์เข้าไป Go จะทำการ copy value
		เอาไว้ และ สร้าง memory address ใหม่
		ตัวอย่าง  func (t *Type) Method() {}  //pointer receiver
				func (t Type) Method() {}   //value receiver
	*/
	//*************Pass by Value**************//
	persons := person{name: "Alice", age: 50}
	changeName(persons)
	fmt.Println("Pass by Value:", persons)

	//*************Pass by Pointer**************//
	persons = person{name: "Alice", age: 50}
	changeNamePPointer(&persons)
	fmt.Println("Pass by Pointer:", persons)
}
