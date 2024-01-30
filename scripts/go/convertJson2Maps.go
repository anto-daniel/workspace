// Go program which converts JSON to maps
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Creating a map using var keyword
	var mymap map[string]string
	// Creating a JSON string
	str := `{"India":"New Delhi","Japan":"Tokyo","France":"Paris","Germany":"Berlin"}`
	// Unmarshalling the JSON string
	json.Unmarshal([]byte(str), &mymap)
	// Printing the map
	fmt.Println("Map after converting JSON")
	fmt.Println(mymap)
}
