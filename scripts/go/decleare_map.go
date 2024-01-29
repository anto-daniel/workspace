// Go program to illustrate how to declare a map

package main

import "fmt"

func main() {
  
    // Creating a map using var keyword
    var mymap map[string]string
  
    // Declaring a map using shorthand declaration
    // mymap := map[string]string
  
    // Using make() function
    // mymap := make(map[string]string)
  
    // Declaring and initializing a map using shorthand declaration
    mymap = map[string]string{
      "India": "Delhi",
      "Japan": "Tokyo",
      "Pakistan": "Islamabad",
      "Australia": "Canberra",
      "UK": "London",
      "Dubai": "UAE",
    }
  
    // Displaying the elements of the map
    for country := range mymap {
      fmt.Println("Capital of", country, "is", mymap[country])
    }
  }

//   Output:
// Capital of India is Delhi
// Capital of Japan is Tokyo
 
// Note: If you try to access a key that does not exist in the map, it will return the zero value of the value type of the map. For example, if you try to access a key that does not exist in the map of type string and string, it will return an empty string. Similarly, if you try to access a key that does not exist in the map of type string and int, it will return 0.
// Example:
 
// package main

// import "fmt"

// func main() {

//     // Creating a map using var keyword
//     var mymap map[string]string

//     // Declaring a map using shorthand declaration
//     // mymap := map[string]string

//     // Using make() function
//     // mymap := make(map[string]string)
 
//     // Declaring and initializing a map using shorthand declaration
//     mymap := map[string]string{
//         "India": "Delhi",
//         "Japan": "Tokyo",
//     }

//     // Displaying the elements of the map
//     for country := range mymap {
//         fmt.Println("Capital of", country, "is", mymap[country])
//     }

//     // Accessing a key that does not exist in the map
//     fmt.Println("Capital of USA is", mymap["USA"])
// }


