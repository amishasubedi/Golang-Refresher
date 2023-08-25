package main

import "fmt"

// constant
const LoginToken string = "Check"

func main() {
	var username string = "Amisha"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isActive bool = false
	fmt.Println(!isActive)
	fmt.Printf("Variable is of type: %T \n", isActive)

	var value uint8 = 255
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	// default values and aliases
	var number int
	fmt.Println("number")
	fmt.Printf("Variable is of type: %T \n", number)

	// implicit type
	var generic = "Kcha"
	fmt.Println(generic)
	fmt.Printf("Variable is of type: %T \n", generic)

	// no var style
	numberOfUser := 3000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	// constants are public
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
