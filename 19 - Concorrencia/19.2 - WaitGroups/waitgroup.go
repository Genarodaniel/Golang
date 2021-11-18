package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//concorrencia != Paralelismo
	//go routine é chamado com o parametro "go" antes da funcao
	// ao usar go routine, o programa não irá esperar o fim da execução da mesma, para seguir a execução

	var waitGroup sync.WaitGroup

	//quantidade de go routines que estarao nesse waitgroup
	waitGroup.Add(2)

	//criação de função anonima que irá chamar as funções
	go func() {
		escrever("Olá mundo")
		//retirando uma função do contador do waitGroup
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em go")
		//retirando uma função do contador do waitGroup
		waitGroup.Done()
	}()

	//o waitGroup wait, irá fazer com que o programa continue executando até o contador de go routines chegar em 0
	// é feito o descréscimo nesse contador, quando é chamado o método waitGroup.Done()

	waitGroup.Wait()
}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
