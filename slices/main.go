package main

import "fmt"

func main() {
	var array1 [10]int = [10]int{1, 2, 3, 4}
	var slice []int = array1[3:]
	fmt.Println(slice)
	fmt.Println("now lets change")
	var slice2 []int = []int{1, 2, 3, 4}
	fmt.Println(cap(slice2))

}
