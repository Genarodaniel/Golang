package main

import "fmt"

// métodos são tipo funções, porém precisam estar associados a uma struct ou a uma interface

type usuario struct {
	nome  string
	idade int8
}

//modo de criar um método que será utilizado por uma struct do tipo usuario

func (u usuario) salvar() {
	//u é a variavel instanciada no momento para chamar os atributos do usuario
	fmt.Println("print dentro do salvar", u.nome)
}

func (u usuario) maior() bool {
	return u.idade >= 18
}

//passar como ponteiro o usuario, para alterar o valor do mesmo
func (u *usuario) fazerAniversario() {
	// não precisa desreferenciar
	u.idade++
}

func main() {

	usuario1 := usuario{"joao", 22}
	fmt.Println(usuario1)
	usuario1.salvar()

	isMaior := usuario1.maior()
	fmt.Println(isMaior)

	fmt.Println(usuario1)
	usuario1.fazerAniversario()
	fmt.Println("---- Fez aniversario --- ")
	fmt.Println(usuario1)
}
