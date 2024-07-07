package main

import "fmt"

func main() {
	var fruits [5]string // the array size ust be declared
	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[3] = "Banana"

	//[2] and [4] are missing which will be indicated as a ' ' (blank) when we will print the array

	fmt.Println("Fruits are: ", fruits)
	fmt.Println("Fruits are: ", len(fruits)) // get the lenth of the array
	// the length of the array is not the values entered but the size of which it is declared

	var vegy = [3]string{"Onion", "potato"} // notice the = sign we have to put in this syntax

	fmt.Println("vegies are: ", vegy)

	var sNo [5]int

	sNo[3] = 45
	fmt.Println(sNo)

}
