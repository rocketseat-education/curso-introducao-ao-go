package main

import (
	"fmt"
)

func main() {
	var pessoas = map[string]string{}
	pessoas["lais"] = 26
	pessoas["leo"] = 32

	if idade, ok := pessoas["lais"]; ok {
		fmt.Println("Pesssoa existe no map", idade, ok)
	} else {
		fmt.Println("Pessoa não existe no map")
	}
	delete(pessoas, "lais")
	fmt.Println(pessoas)
}
