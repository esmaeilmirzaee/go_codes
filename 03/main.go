package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#f00",
		"green": "#0f0",
	}
	colours["white"] = "#fff"
	fmt.Println(colours)
	printMap(colours)
}

func printMap(colours map[string]string) {
	for c, h := range colours {
		fmt.Printf("%v is %v", c, h)
	}
}
