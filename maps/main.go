package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {

	// ways of declaring map
	var codes map[string]string
	fmt.Println("Codes Map : ", codes) // gets initialized with empty value unlike Java

	codes2 := make(map[string]person)    // value is of type person which is struct
	fmt.Println("Codes2 Map : ", codes2) // gets initialized with empty value

	colors := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}
	printMap(colors)

	//manipulate maps
	colors["Z"] = 26
	delete(colors, "D")
	fmt.Println("\nPrinting colors after manipulation :")
	printMap(colors)
}

func printMap(c map[string]int) {
	for key, value := range c {
		fmt.Println("Key : ", key, ", Value : ", value)
	}
}
