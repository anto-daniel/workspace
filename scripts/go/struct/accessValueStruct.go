// Golang program to access fields of struct
package main

import "fmt"

// define struct

type bike struct {
	name, model, color string
	weight_in_kg       float64
}

func main() {
	b := bike{name: "Pulsar", model: "220F", color: "Black", weight_in_kg: 160}

	// Access struct fields using the dot operator
	fmt.Println("bike name:", b.name)
	fmt.Println("bike color:", b.color)

	// Assign a new value to a struct field
	b.color = "Neon Red"
	fmt.Println("bike:", b)
}
