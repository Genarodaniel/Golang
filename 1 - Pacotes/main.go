package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	auxiliar.Escrever2()
	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)

}
