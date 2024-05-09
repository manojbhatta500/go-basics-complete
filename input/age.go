package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("You were born in  :")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("\nyou will be %d years old at the end of 2024", 2024-input)

}
