// Wite a Go Program to add key value pairs

package main

import "fmt"

func main() {
  var my_map = map[int]string{
    1: "US",
    91: "India",
    86: "China",
    44: "UK",
  }
  fmt.Println("Original Map\n", my_map)

  // Adding new key-value pairs in the map
  my_map[39] = "Italy"
  my_map[55] = "Brazil"
  fmt.Println("\nMap after adding new key-value pairs\n", my_map)
}





