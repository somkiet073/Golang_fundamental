package employee

import "fmt"

type employee struct {
	Firstname   string
	Lastname    string
	TotalLeaves int
	LeavesToken int
}

var employeeInstance *employee

func Init(
	firstname string,
	lastname string,
	totalLeaves int,
	leavesToken int,
) *employee {

	// nil = ค่าว่าง
	if employeeInstance == nil {
		employeeInstance = &employee{
			Firstname:   firstname,
			Lastname:    lastname,
			TotalLeaves: totalLeaves,
			LeavesToken: leavesToken,
		}
	}

	return employeeInstance
}

func (e employee) LeavesRemaining() { // การ add func เข้า struct โดยประกาศ func ตามด้วย (ตัวแปร ชื่อ Struct)
	fmt.Printf("%s  %s  has %d leaves remaining:\n", e.Firstname, e.Lastname, e.TotalLeaves)
}
