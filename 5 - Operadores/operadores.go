package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 1 / 2
	multiplicacao := 1 * 2
	resto := 1 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	// não é possível fazer operações matematicas, nem comparacoes entre tipos diferentes no GO (int32 e int 16 exemplo)
	var numero1 int16 = 10
	var numero2 int16 = 20
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM aritmeticos

	//ATRIBUIÇÕES

}
