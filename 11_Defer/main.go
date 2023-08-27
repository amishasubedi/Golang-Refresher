package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start of main function")

	// Create a new file called sample.txt
	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}

	// Ensure the file is closed at the end of this function using defer.
	// This means even if we encounter an error while writing to the file, it will still be closed properly.
	defer file.Close()

	// Writing some lines to the file
	_, err = file.WriteString("Hello, Go!\n")
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	// Another deferred function. Note: Deferred functions are executed in LIFO (Last-In-First-Out) order.
	// This means this line will print before the file closes, but after any further operations in main.
	defer fmt.Println("File write operation completed (whether successful or not).")

	fmt.Println("End of main function")
}
