package main

import "fmt"

func main() {
	var arr1 = [2]int{1, 2}
	arr2 := [7]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("this is another start")

	var arr3 = [10]int{1, 2, 3, 4, 5}

	fmt.Println(arr3)
	fmt.Println("now the array after changing index")

	arr3[0] = 5

	fmt.Println(arr3)

	fmt.Println("the length of array 1 is:", len(arr1))

	fmt.Println("the length of array 2 is:", len(arr2))

	fmt.Println("the length of array 3 is:", len(arr3))

}
