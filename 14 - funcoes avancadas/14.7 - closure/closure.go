package main

import "fmt"

func closure() func() {
	text := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func main() {

	texto := "Dentro da main"
	fmt.Println(texto)

	novafuncao := closure()
	novafuncao()

}
