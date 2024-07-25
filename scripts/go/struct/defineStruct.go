package main

import "fmt"

// Define a struct type
type city_list struct {
	city    string
	state   string
	country string
}

func main() {
	// Declare and initialize struct using a struct literal
	c1 := city_list{"Chennai", "Tamil Nadu", "India"}
	c2 := city_list{"Tokyo", "Tokyo", "Japan"}
	c3 := city_list{"Paris", "Ile-de-France", "France"}

	fmt.Println("City 1: ", c1)
	fmt.Println("City 2: ", c2)
	fmt.Println("City 3: ", c3)
}
