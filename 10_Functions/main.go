package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()

	num := 40

	if isEven(num) {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	proResult := proAdder(2, 5, 6, 7)
	fmt.Println(proResult)

}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total = total + val
	}

	return total
}

// Function to check if a number is even
func isEven(n int) bool {
	return n%2 == 0
}

func greeter() {
	fmt.Println("Namaste!")
}
