package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number: ")

	input, _ := reader.ReadString('\n') //remember that the string will include \n (a line break) in it

	fmt.Println("input string is : ", input)

	// add 1 to the input number which is in string format

	/*
		number, err := strconv.ParseInt(input, 10, 32)


		// the input contains a '\n' at the end hence cant be parsed into int
		// hence above parsing will give error
		if err != nil {
			fmt.Println("error occured", err)
		} else {
			number = number + 1
			fmt.Println(number)
		}
	*/

	number, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	// if error handling is not needed just put an _ there

	number = number + 1
	fmt.Println("new number is: ", number)

}
