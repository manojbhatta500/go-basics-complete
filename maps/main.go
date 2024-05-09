package main

import "fmt"

func main() {

	var map1 map[string]int = map[string]int{
		"bugati": 1000,
		"honda":  2000,
		"lauda":  3000,
	}

	fmt.Println(map1)

	fmt.Println(map1["bugati"])

}
