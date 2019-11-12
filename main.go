package main

import "fmt"

func main() {
	// declaring a string to string map
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}
	fmt.Println(colors)

	// square brace syntax
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// deleting call where 1st arg is map and 2nd is key in the map
	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
