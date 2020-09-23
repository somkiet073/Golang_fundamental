package main

import "fmt"

func main() {
	i := 1
	// like a while loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// while loop break
	for {
		fmt.Println("loop")
		break
	}

	// for loop 
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
