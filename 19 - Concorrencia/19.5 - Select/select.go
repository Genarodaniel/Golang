package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		//select pega a mensagem que estiver pronta pra ser recebida em cada loop, e faz alguma ação com ela, se ambas estiverem prontas faz com as duas, se não, pega somente a que estiver pronta
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

		// messageCanal1 := <-canal1
		// fmt.Println(messageCanal1)

		// messageCanal2 := <-canal2
		// fmt.Println(messageCanal2)
	}
}
