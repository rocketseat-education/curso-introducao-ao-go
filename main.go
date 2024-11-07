package main

import "fmt" //check


func main() {
	
	users := map[string]string{
		"nome": "Joao",
		"idade": "30",
	}

	for key, _ := range users {
		fmt.Println((key))
	}

}
