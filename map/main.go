package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"blue": "5c3424",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	colors["white"] = "#ffff"
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c{
		fmt.Println("hex code for", color, "-", hex)
	}
}