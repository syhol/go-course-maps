package main

import (
	"fmt"
)

func main() {
	// examples()

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}

func examples() {
	var colors1 map[string]string
	fmt.Println(colors1)

	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"
	fmt.Println(colors2)

	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	delete(colors3, "red")
	fmt.Println(colors3)
}
