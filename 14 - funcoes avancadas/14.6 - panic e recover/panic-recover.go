package main

import "fmt"

func recuperarExecucao() {
	// verificando se o programa conseguiu ser recuperado, e caso sim printando na tela
	// não há problema em chamar o recover, caso não seja necesário
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso")
	}
}
func aprovado(n1, n2 float64) bool {

	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// panic é utilizado quando não há nenhum return tratando a condição que foi gerada de acordo com os parametros
	// ela interrompe a execução da função e chama todas as funções que tem defer
	// é necessário utilizar recover para voltar a execução
	// como são chamadas todas as funções com defer na hora que entra o panic, a função de recuperação tem que ser chamada com defer

	panic("A média é exatamente 6")
}

func main() {

	aprovado(6, 6)
	fmt.Println("pós execução")

}
