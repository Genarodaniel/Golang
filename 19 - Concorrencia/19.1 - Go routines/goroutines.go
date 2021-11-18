package main

import (
	"fmt"
	"time"
)

func main() {

	//concorrencia != Paralelismo
	//go routine é chamado com o parametro "go" antes da funcao
	// ao usar go routine, o programa não irá esperar o fim da execução da mesma, para seguir a execução
	go escrever("Olá mundo!")
	escrever("Programando em go")
}

func escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
