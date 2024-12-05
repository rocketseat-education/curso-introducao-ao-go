package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
}

func (p *Pessoa) Apresentar() {
	p.Nome = "Joana"
	fmt.Printf("Olá, meu nome é %s e tenho %d anos. \n", p.Nome, p.Idade)
}

func main (){
	p1 := Pessoa{Nome: "Lais", Idade:26}
	p1.Apresentar()
	fmt.Println(p1.Nome)
}
