package main

import "fmt"

func main() {
	var a []int = []int{100, 200, 300, 4000, 500}

	for i, element := range a {
		fmt.Printf("%v : %v \n", i, element)
	}
}
