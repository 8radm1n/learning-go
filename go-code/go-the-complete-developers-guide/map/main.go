package main

import "fmt"

func main() {
	// Create a map with keys of the type string
	// and values of the type string
	// all the keys and values must be of the
	// same type declared when creating the map.
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// Alternate syntax to create a map
	// with zero values assigned.
	// var colours map[string]string
	// or
	// colours := make(map[string]string)

	// Add values to a map
	colours["white"] = "#ffffff"
	colours["black"] = "#000000"

	// Delete keys from a map
	// delete(colours, "white")

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Printf("%v - %v \n", colour, hex)
	}
}
