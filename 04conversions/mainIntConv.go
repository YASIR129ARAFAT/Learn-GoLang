// Golang program to convert String
// into Integer Data type
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	str := "GeeksforGeeks"

	fmt.Println("Before :", reflect.TypeOf(str))
	fmt.Println("String value is: ", str)
	i, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println("After :", reflect.TypeOf(i))
	fmt.Println("Integer value is: ", i, "\n")

	// parsing the string "100" to int in decimal system (base 10)

	str1 := "100"

	fmt.Println("Before :", reflect.TypeOf(str1))
	fmt.Println("String value is: ", str1)
	i1, _ := strconv.ParseInt(str1, 10, 64) // the base is given as 10 and size for int is 64 bits
	fmt.Println("After :", reflect.TypeOf(i1))
	fmt.Println("Integer value is: ", i1, "\n")

	// parsing the string "1011" to int in binary system (base 2)
	str2 := "1011"

	fmt.Println("Before :", reflect.TypeOf(str2))
	fmt.Println("String value is: ", str2)
	i2, _ := strconv.ParseInt(str2, 2, 64) // the base is given as 2 and size for int is 64 bits
	fmt.Println("After :", reflect.TypeOf(i2))
	fmt.Println("Integer value is: ", i2, "\n")
}
