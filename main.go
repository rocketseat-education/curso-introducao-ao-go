package main

import "fmt"

// int8: O intervalo de valores é de -128 a 127. 
// int16: O intervalo de valores é de -32.768 a 32.767. 
// int32: O intervalo de valores é de -2.147.483.648 a 2.147.483.647. 
// int64: O intervalo de valores é de -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807.

func main() {
	var idade int = 30
	var contador int32 = 2
	var indice int8 = 1

	fmt.Println("Idade:", idade )
}
