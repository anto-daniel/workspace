// Go program to check if a key exists in a map
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
	fmt.Println("Map before checking")
	fmt.Println(mymap)
	// Checking if the key "India" exists in the map
	value, ok := mymap["India"]
	// If the key exists, print the value
	if ok == true {
		fmt.Println("Value of key India is", value)
	} else {
		fmt.Println("Key does not exist")
	}
}
