package main

import "fmt" //check

type Cliente struct {
	Nome string
	Idade int
	Endereco Endereco
	Email string
}


type Endereco struct {
	Rua string
	Numero int
	Cep string
	Estado string
}

func main() {
	
	cliente1 := Cliente{
		Nome: "lais",
		Idade: 26,
		Endereco: Endereco{
			Rua: "Rua das Flores",
			Numero: 123,
			Estado: "SP",
		},
	}

	cliente2 := Cliente{
		Nome: "juan",
		Idade: 39,
	}

	cliente2.Email = "juan@google.com"

	fmt.Println(cliente1.Endereco.Numero)
	cliente1.Endereco.Numero = 124
	fmt.Println(cliente1.Endereco.Numero)

	fmt.Println(cliente2)
}
