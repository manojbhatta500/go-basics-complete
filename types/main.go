package main

import "fmt"

func main() {

	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const pi = 3.14122

	fmt.Println(pi)

	const typeConstant int = 2

	fmt.Println(typeConstant)

}
