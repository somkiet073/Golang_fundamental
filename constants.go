package main

// เรียกใช้งาน Package fmt, math
import (
	"fmt"
	"math"
)

/*
	Constant คือการประกาศค่าคงที่ 
	คำสั่ง const นั่นสามารถทำงานที่ไหนก็ได้ที่คำสั่ง var ทำงานได้
	
*/

// ประกาศตัวแปรคงที่ s Type string
const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
