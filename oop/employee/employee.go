package employee

import (
	"fmt"
)

type Employee struct { // การประกาศชื่อ Struct ตัวใหญ่เพื่อให้ภายนอกสามารถเรียกใช้งานได้
	Firstname   string
	Lastname    string
	TotalLeaves int
	LeavesToken int
}

func (e Employee) LeavesRemaining() { // การ add func เข้า struct โดยประกาศ func ตามด้วย (ตัวแปร ชื่อ Struct)
	fmt.Printf("%s  %s  has %d leaves remaining:\n", e.Firstname, e.Lastname, e.TotalLeaves)
}
