package main

import "fmt"

func main() {

	// first just declare array   and print it

	var arr1 [5]int

	var arr2 [3]float64

	fmt.Println(arr1)

	fmt.Println(arr2)

	// do some calculation using loop and array maybe add and remove
	ulta := 100
	fmt.Println("the total length of ulta is ", len(arr1))
	for i := 0; i <= len(arr1)-1; i++ {

		arr1[i] = ulta
		ulta--
	}

	fmt.Println(arr1)

	fmt.Println(arr2)

}
