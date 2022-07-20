package main

import "fmt"

func main() {

	//Another way of declaring maps
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//	colors["White"] = "#ffffff"

	//	delete(colors, "White")

	//	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
