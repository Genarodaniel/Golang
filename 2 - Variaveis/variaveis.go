package main

import "fmt"

func main() {

	//declaracao com tipo explicito
	var teste string = "testando"

	//declaracao com tipo implicito
	teste2 := "testando 2"
	fmt.Println(teste)
	fmt.Println(teste2)

	//declarar mais de uma variavel com tipo explicito
	var (
		teste3 string = "teste3"
		teste4 string = "teste4"
	)

	fmt.Println(teste3, teste4)

	//declarar mais de uma variavel sem tipo explicito
	teste5, teste6 := "teste 5", "teste 6"

	fmt.Println(teste5, teste6)

	//declarar constante (valor não pode mudar)
	const constante1 string = "churrasco de constante"
	fmt.Println(constante1)

	//trocar valor de variavel
	teste5, teste6 = teste6, teste5
	fmt.Println(teste5, teste6)

}
