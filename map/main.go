package main

import (
	"fmt"
)

func main() {
	// var colours map[string]string

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4b4753",
		"white": "#ffffff",
	}

	// colours := make(map[string]string)
	// colours["white"] = "#ffffff"

	// delete(colours, "white")

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println(colour, hex)
	}
}
