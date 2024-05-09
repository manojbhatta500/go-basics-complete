package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	print("please enter your age : ")
	scanner.Scan()

	inputString := scanner.Text()

	input, err := strconv.ParseInt(inputString, 10, 64)

	if err != nil {
		fmt.Println("Input must me Integer")
		return
	}

	if input > 18 {
		fmt.Println("you can ride")
	} else if input > 14 {
		fmt.Println("you can ride with your parent")
	} else {
		fmt.Println("you can't ride.")
	}

}
