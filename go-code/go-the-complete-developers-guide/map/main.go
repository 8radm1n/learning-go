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
	fmt.Println(colours)
}
