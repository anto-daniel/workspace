package main

import "fmt"

type employee struct {
	name string
	age  int
	city string
}

func main() {
	var emp1 employee // zero valued struct
	fmt.Println("Employee 1 details: ", emp1)
}
