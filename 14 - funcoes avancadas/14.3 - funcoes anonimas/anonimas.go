package main

import "fmt"

func main() {

	//declara a função anonima e os "()" dps da declaração, quer dizer que assim que declarada a mesma já deve ser executada
	//dentro do "()" final, é possível passar os parametros de execução
	// precisa definir o tipo de retorno, caso for retornar
	recebido := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("parametro texto")

	fmt.Println(recebido)
}
