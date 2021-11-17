package main

import (
	"fmt"
	"time"
)

func main() {

	//for comum
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println("incrementando i")
		i++

	}
	// fmt.Println(i)

	//for com variavel sendo inicializada diretamente (for init)

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando J", j)
		time.Sleep(time.Second)
	}

	//for com clausula range, que serve para iterar arrays e slices

	nomes := []string{"João", "Pedro", "Carlos"}

	// primeiro parametro sempre o indice, o segundo é sempre o valor
	for indice, nome := range nomes {
		fmt.Println("indice: ", indice)
		fmt.Println("valor:", nome)
	}

	// é possível não utilizar o indice, colocando o "_" como parâmetro
	for _, nome := range nomes {
		fmt.Println("valor:", nome)
	}

	// é possível não utilizar o indice, colocando o "_" como parâmetro
	for indice, letra := range "PALAVRA" {
		fmt.Println("indice:", indice)
		fmt.Println("valor:", letra)
		// use a função string para pegar o valor em string do código da tabela asc retornado
		fmt.Println("Valor em string:", string(letra))
	}

	usuario := map[string]string{
		"nome":      "teste",
		"sobrenome": "testando",
	}

	// iterar um MAP
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// não é possível iterar struct

}
