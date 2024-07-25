package main

import "fmt"

// Define a struct type
type city_list struct {
	city    string
	state   string
	country string
}

func main() {
	// Naming fields while initializing a struct
	c1 := city_list{city: "Chennai", state: "Tamil Nadu", country: "India"}
	c2 := city_list{city: "Tokyo", state: "Tokyo", country: "Japan"}
	c3 := city_list{city: "Paris", state: "Ile-de-France", country: "France"}

	fmt.Println("City 1: ", c1)
	fmt.Println("City 2: ", c2)
	fmt.Println("City 3: ", c3)
}
