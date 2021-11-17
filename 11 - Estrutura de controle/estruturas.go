package main

import "fmt"

func main() {

	// if comum
	numero := 10

	if numero > 15 {
		fmt.Println("maior que a 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	//setando variaveis dentro do próprio if (ifinit)
	// ao criar uma variavel no if init, ela só é acessada dentro do if, não pode ser acessada por fora
	if outronumero := numero; outronumero > 0 {
		fmt.Println("é maior que 0")
	} else {
		fmt.Println("é menor que 0")
	}

}
