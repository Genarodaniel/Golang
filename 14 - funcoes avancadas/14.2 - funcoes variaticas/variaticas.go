package main

import "fmt"

//criando função que recebera a quantidade parametros variavel
//ao usar os "..." definimos que a quantidade de parametros podera ser de 0 até n
func soma(numeros ...int) (total int) {
	for _, numero := range numeros {
		total += numero
	}
	return
}

// função com um parametro variavel e outro parametro fixo
// só é possível ter um parametro variavel por função, e este deve obrigatoriamente ser o ultimo
func escrever(text string, numeros ...int) {

	for _, numero := range numeros {
		fmt.Println(text, numero)
	}
}

func main() {
	result := soma(1, 2, 3, 4, 5, 5, 6)
	fmt.Println(result)

	escrever("numero:", 1, 3, 4, 5, 5, 6)
}
