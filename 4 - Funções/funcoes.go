package main

import "fmt"

func Somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := Somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("escrevendo na tela")
	fmt.Println(resultado)

	// funcoes que retornam mais de um valor como retorno.
	soma, subtracao := calculos(10, 5)
	fmt.Println(soma, subtracao)

	// ao usar o '_' ao invés de atribuir o valor para uma variavel, queremos dizer que não utilizaremos aquele retorno
	resSoma, _ := calculos(10, 5)
	fmt.Println(resSoma)
}
