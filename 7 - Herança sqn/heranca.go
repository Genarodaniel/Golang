package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	// ao passar outra struct sem referenciar o tipo, é a mesma coisa que somente copiar os campos de uma para outra
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"João", "Pedro", 20, 198}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
}
