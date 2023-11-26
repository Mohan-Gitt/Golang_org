package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string  -- empty value colors[]

	// colors := make(map[int]string)
	// colors[10] = "fffff"
	// colors[20] = "gh065"

	// delete(colors, 10)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4bf075",
		"white": "562bj",
	}

	// fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
