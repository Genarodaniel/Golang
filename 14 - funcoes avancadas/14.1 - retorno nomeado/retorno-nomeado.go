package main

import "fmt"

//função com retorno nomeado, não precisa redeclarar a variavel, o mesmo que foi definido no tipo de retorno, sera retornado no fim da função
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {

	soma, subtracao := calculosMatematicos(3, 5)
	fmt.Println(soma, subtracao)
}
