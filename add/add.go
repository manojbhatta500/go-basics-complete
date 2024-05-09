package main

import (
	"fmt"
)

func main() {

	x := take(2, 5)

	fmt.Println(x)

}

func take(a, b int) int {
	return a + b
}
