package main

import (
	"fmt"
)

func main() {
	var gavetas []string
	gavetas = append(gavetas, "copos", "panos", "pratos")
	gavetas = gavetas[:2]
	fmt.Println(gavetas)
}
