package main

import "fmt"

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Color:", color, "Hex:", hex)
	}
}
