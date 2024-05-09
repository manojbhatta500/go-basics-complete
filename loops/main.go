package main

import "fmt"

func main() {

	// x := 0

	// fmt.Println("this is first type of loop")

	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// fmt.Println("this is second type of loop")

	// for x := 0; x <= 10; x++ {
	// 	fmt.Println(x)
	// }
	for x := 2; x <= 100; x++ {
		if x%2 == 0 {
			fmt.Println(x, "is even number")
		} else {
			fmt.Println(x, "is an odd number")
		}
	}

}
