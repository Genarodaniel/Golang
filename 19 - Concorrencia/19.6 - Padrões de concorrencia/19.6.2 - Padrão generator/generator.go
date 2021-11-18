package main

import (
	"fmt"
	"time"
)

func main() {
	// para chamar o canal não precisa do go

	canal := escrever("Olá mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

//importante salientar que o canal retornado é um canal do tipo apenas de recebimento <- antes do chan

//padrão generator encapsula a chamada da go routine e tira da main
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
