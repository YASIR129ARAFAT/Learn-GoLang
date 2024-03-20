// Golang program to convert
// String into Float Data type
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	str := "3.1415926535"

	fmt.Println("Before :", reflect.TypeOf(str))
	fmt.Println("String value is: ", str)
	f, _ := strconv.ParseFloat(str, 64) // the result will be float64 type
	fmt.Println("After :", reflect.TypeOf(f))
	fmt.Println("Float value is: ", f)

	str1 := "3.1415926535"

	fmt.Println("Before :", reflect.TypeOf(str1))
	fmt.Println("String value is: ", str1)
	f1, _ := strconv.ParseFloat(str1, 32) // the result will be still of float64 type
	fmt.Println("After :", reflect.TypeOf(f1))
	fmt.Println("Float value is: ", f1, "\n")

}
