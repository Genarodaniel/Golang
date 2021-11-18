package main

import "fmt"

func main() {

	// o segundo parametro especifica o tamanho do canal
	// o canal com buffer só bloqueia quando atingir a quantidade máxima(2)
	canal := make(chan string, 2)

	canal <- "Olá mundo"
	canal <- "Programando em go"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
