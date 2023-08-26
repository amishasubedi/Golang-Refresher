package main

import (
	"fmt"
	"sort"
)

func main() {

	//******************ARRAY*********************
	fmt.Println("Welcome to Array")

	var numbers [4]int

	numbers[0] = 1
	numbers[1] = 2
	numbers[3] = 3

	fmt.Println("Numbers List is: ", numbers)
	fmt.Println("Numbers List is: ", len(numbers)) // prints 4 because of the reserved memory

	var fruits = [3]string{"Apple", "Oranges", "Banana"}
	fmt.Println("Fruits List is: ", fruits)

	//****************SLICES*************************
	// Slices is used for in more golang - Slices automatically expands memory unlike arrays
	fmt.Println("Welcome to Slices")

	var numberList = []int{2, 3, 4, 5, 6, 7}
	fmt.Printf("Type of number list: %T\n", numberList)
	fmt.Println("Initial number list: ", numberList)

	// adding in slices
	numberList = append(numberList, 8)
	fmt.Println("Number List after appending: ", numberList)

	numberList = append(numberList[3:])                  // slice up
	fmt.Println("Number List after slice: ", numberList) // starts from position 3

	numberList = append(numberList[1:3])                 // slice up
	fmt.Println("Number List after slice: ", numberList) // starts from position 1, 3 is non-inclusive

	// allocate memory using make
	names := make([]string, 2)
	names[0] = "Ram"
	names[1] = "Shyam"

	names = append(names, "Hari", "Alu") // append reallocate memory to accomodate all values
	fmt.Println("Name List: ", names)

	// sorting
	sort.Ints(numberList)
	fmt.Println("Sorted number list", numberList)

	fmt.Println(sort.IntsAreSorted(numberList))

	// removing values from slice based on index
	var courses = []string{"React", "Vue", "Javascript", "Python"}
	fmt.Println("Courses: ", courses)

	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Using append to remove", courses)

}
