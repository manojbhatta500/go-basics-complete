package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your valid number: ")
	scanner.Scan()
	input := scanner.Text()

	realInput, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Println(" sorry can't use the text")
		return
	}

	switch realInput {

	case 1:
		fmt.Println("sunday")

	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("tuesday")
	case 4:
		fmt.Println("wednesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")

	case 7:
		fmt.Println("saturday")

	default:
		fmt.Println("sorry invalid number")

	}

}
