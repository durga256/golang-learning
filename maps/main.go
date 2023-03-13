package main

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b5645",
	}

	//or
	//colors2 := make(map[string]string)
	//to assign
	colors["white"] = "$ffff"

	//delete(colors, "white")
	printMap(colors)
}
