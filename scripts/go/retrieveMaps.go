// Go program to retrieve key-value pair from a map and print the retrieved value
package main

import "fmt"

func main() {
	// Creating a map using var keyword
	var mymap map[string]string
	// Adding key-value pairs to the map
	mymap = map[string]string{
		"India":   "New Delhi",
		"Japan":   "Tokyo",
		"France":  "Paris",
		"Germany": "Berlin",
	}
	// Printing the map
	fmt.Println("Map before retrieving")
	fmt.Println(mymap)
	// Retrieving the value of key "India"
	value := mymap["India"]
	// Printing the value
	fmt.Println("Value of key India is", value)
}
