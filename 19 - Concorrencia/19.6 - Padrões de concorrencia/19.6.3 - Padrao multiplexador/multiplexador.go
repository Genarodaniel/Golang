package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(escrever("Olá mundo"), escrever("Programar em go"))

	for i := 0; i < 50; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canaldeEntrada1, canaldeEntrada2 <-chan string) <-chan string {

	canalDeSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canaldeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canaldeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
