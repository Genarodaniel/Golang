package main

import "fmt"

// funcao recursiva para informar o numero que esta na posicao passada como parametro na sequencia de fibonacci
// funcoes recursivas precisam ter a informação de quando a mesma irá parar de ser executada
func fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	posicao := uint(10)

	//for para iterar a fibbonaci ate o resultado
	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
