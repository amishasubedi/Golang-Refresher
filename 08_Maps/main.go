package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	courses := make(map[string]string)
	courses["JS"] = "JavaScript"
	courses["RB"] = "Ruby"
	courses["PY"] = "Python"

	fmt.Println("List of all courses: ", courses) // not comma sepearted
	fmt.Println("JS shorts for: ", courses["JS"]) // key, value

	// delete by key
	delete(courses, "RB")
	fmt.Println("List of all courses: ", courses)

	// loop through maps
	for key, value := range courses {
		fmt.Println("For key ", key, "value is ", value)
	}

}
