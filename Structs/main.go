package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Structs in Golang")

	data := User{"Amisha", "amisha@gmail.com", true, 20}
	fmt.Println(data)
	fmt.Printf("Details are: %+v\n ", data)

	data.GetStatus()
	data.NewMail()

	// if else
	if data.Age > 21 {
		fmt.Println("Adult")
	} else if data.Age > 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}

	// switch case
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice is 1")
	case 2:
		fmt.Println("Dice is 2")
		fallthrough
	case 3:
		fmt.Println("Dice is 3")
		fallthrough
	case 4:
		fmt.Println("Dice is 4")
		fallthrough
	case 5:
		fmt.Println("Dice is 5")
	case 6:
		fmt.Println("Dice is 6")
	default:
		fmt.Println("Default")
	}

	// loops
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// simple for loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Println(i, day)
	// }

	value := 1

	// break and continue
	for value < 10 {

		if value == 2 {
			goto lco
		}

		if value == 9 {
			break
		}

		if value == 2 {
			value++
			continue
		}

		fmt.Println(value)
		value++
	}

lco:
	fmt.Println("Goto Demo")

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// methods
func (u User) GetStatus() {
	fmt.Println("Status is: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test.gmail.com"
	fmt.Println("New Email is: ", u.Email)
}
