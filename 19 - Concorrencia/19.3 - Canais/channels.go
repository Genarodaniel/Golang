package main

import (
	"fmt"
)

func main() {

	// é possível deixar diversas go routines rodando ao mesmo tempo, e esperar o valor de alguma com a notação de recebimento de valor do canal: (mensagem := <- canal), essa clausula, faz com que seja esperado um retorno de valor para o canal
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	//notação que tem <- para receber
	for {
		//primeira variavel é valor que o canal recebe, segunda é se o canal ainda está aberto
		//evita deadlocks
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// notação que não precisa verificar se o canal esta aberto ou fechado
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//envia a informação pro canal
		canal <- texto
	}

	//fecha o canal
	close(canal)
}
