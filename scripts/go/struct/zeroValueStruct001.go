package main

import "fmt"

type employee struct {
	name     string
	age      int
	location string
}

func main() {
	emp1 := employee{
		age:      30,
		location: "Pune",
	}
	fmt.Println("Employee 1 details: ", emp1)
}
