package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Handling in Go")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) // 01-02-2006 is a syntax in go documentation for date formatting
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.June, 30, 10, 3, 3, 7, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
