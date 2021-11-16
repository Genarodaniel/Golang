package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int = 100
	var ponteiro *int

	ponteiro = &variavel3
	// colocar o & faz com que o valor a ser referenciado seja o endereço de memória
	// colocar o * antes do ponteiro, faz com que o valor a ser utilizado sejá o valor que esta guardado dentro daquele endereço de memória
	fmt.Println(variavel3, *ponteiro) // desreferênciação

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
