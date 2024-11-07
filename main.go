package main

import (
	"fmt"
)

func main() {
	
	players := map[string]int{
		"lais": 26,
	}

	if value, ok := players["eduardo"]; ok {
		fmt.Println("pontos:", value, ok)
	}
}
