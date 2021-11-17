package main

import "fmt"

func funcao1() {
	fmt.Println("executando funcao 1")
}

func funcao2() {
	fmt.Println("executando funcao 2")
}

func aprovado(n1, n2 float32) bool {

	defer fmt.Println("Media calculada, resultado retornado", n1)
	fmt.Println("Entrando na função para verificar se esta aprovado ou não", n2)
	media := (n1 + n2) / 2

	if media >= 6 {

		return true
	}

	return false
}

func main() {
	//defer = adiar, ao utilizar essa clausula, a execução da função será adiada até o ultimo momento possível

	// defer funcao1()
	// funcao2()

	// o defer deve ser utilizado na execução.
	// exemplo, ao utilizar defer no fmt.println, e chamar uma função dentro, ela será executada em ordem, porém o print será exibido depois.
	// já ao utilizar defer na invocação da função, a função será chamada por ultimo

	//defer pode ser utilizado principalmente para fechar a conexao com o banco de dados, assim garantimos que a mesma será fechada ao fim da execução do bloco de códigos que chama o banco
	defer aprovado(6, 3)
	aprovado(5, 7)
	// fmt.Println(aprovado(6, 3))
	// fmt.Println(aprovado(5, 7))
}
