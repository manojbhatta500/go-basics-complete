package main

import "fmt"

func main() {
	add, sub := normalInt(1000, 100)
	fmt.Println("addition", add)

	fmt.Println("subtraction", sub)

}

func normalInt(a, b int) (int, int) {
	return a + b, a - b
}
