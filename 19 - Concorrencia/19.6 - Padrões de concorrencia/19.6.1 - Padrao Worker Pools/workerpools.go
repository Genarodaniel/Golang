package main

import "fmt"

func main() {

	//canal de numeros a serem calculados
	tarefas := make(chan int, 100)
	//canal de resultados armazenados
	resultados := make(chan int, 100)

	// ao chamar o canal, ele ira esperar ter algum valor no mesmo para iniciar.
	//no momento dessa chamada não há nada
	//cada vez que go é chamado, será separado um processo para execução.
	// em prática, cada chamada dobra a velocidade
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 100; i++ {
		tarefas <- i
	}
	//após o for, o canal foi populado e será executado na medida que uma mensagem é inserida no mesmo

	//fecha o canal
	close(tarefas)

	//para imprimir

	for mensagem := range resultados {
		fmt.Println(mensagem)
	}

}

// é possível especificar se o canal irá só receber ou só enviar dados
// <- antes do chan para só receber
// <- depois do chan para canais que só enviam
func worker(tarefas chan int, resultados chan int) {

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

	close(resultados)

}

func fibonacci(posicao int) int {

	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
