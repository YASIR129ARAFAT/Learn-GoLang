package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi there")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("you entered ", name)

	// var name = reader.ReadString("\n")
}
