package main

import "fmt"

/*
UNDERSTANDING POINTERS IN GENERAL

Why pointers exist?
	- Pointers exist to provide direct access to memory locations. Instead of passing copies of variables to functions,
	  which can lead to inconsistencies, pointers allow us to pass the actual memory reference, ensuring the
	  original data is accessed and modified. This avoids unnecessary duplication and potential discrepancies.

*/

func main() {
	fmt.Println("Welcome to Pointers")

	// pointer storing integer value
	var ptr *int

	fmt.Println("Value of Pointer is: ", ptr) // <nil>

	myNumber := 23

	// create another pointer to reference to myNumber
	var anotherPtr = &myNumber
	fmt.Println("Another Pointer: ", anotherPtr)  // prints memory address
	fmt.Println("Another Pointer: ", *anotherPtr) // prints the value inside the memory address

	*anotherPtr = *anotherPtr + 2
	fmt.Println("New value: ", anotherPtr) // creates another memory address to store new value
	fmt.Println("New Value: ", *anotherPtr)

}
