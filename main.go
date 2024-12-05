package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main (){
	var p1 Pessoa = Pessoa{Nome: "lais"}
	var p3 *Pessoa = &p1

	p3.Nome = "Vanessa"

	fmt.Println(p1)
	fmt.Println(p3)

}
