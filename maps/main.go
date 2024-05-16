package main

import "fmt"

func main(){
	// This is a map.
	// A map is a collection of key-value pairs.
	// The keys must be of the same type and the values must be of the same type.
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"blue": "#0000ff",
	}
	// var colors map[string]string // This is also valid
	// colors := make(map[string]string) // This is also valid

	fmt.Println(colors)

	// Adding a new key-value pair
	colors["white"] = "#ffffff"
	fmt.Println("Adding a color: white")
	fmt.Println(colors)

	// Deleting a key-value pair
	delete(colors, "white")
	fmt.Println("Deleting a color: white")
	fmt.Println(colors)

	// Using the printMap function
	printMap(colors)
}

func printMap(c map[string]string) {
	// Iterating over a map
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}