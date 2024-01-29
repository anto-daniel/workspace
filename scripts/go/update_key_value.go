// Go Program to update  key-value pair in a map and print the updated map
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
  fmt.Println("Map before updating")
  fmt.Println(mymap)
  // Updating the value of key "India"
  mymap["India"] = "Chennai"
  // Printing the map
  fmt.Println("Map after updating")
  fmt.Println(mymap)
}
