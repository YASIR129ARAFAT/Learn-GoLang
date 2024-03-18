package main

import "fmt"

const LoginToken string = "ghabbhhjd" // Public

// id:="alpha@iebc" // walrus operator declaration is not allowed in go lang

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "abcdef.in"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	const PI = 3.14159 // Untyped float constant
	fmt.Println(PI)
	fmt.Printf("Variable is of type: %T \n", PI)

	const MAX_USERS int = 100 // Typed integer constant
	fmt.Println(MAX_USERS)
	fmt.Printf("Variable is of type: %T \n", MAX_USERS)

}
