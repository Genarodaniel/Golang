package main

import "fmt"

//ao passar uma interface dessa forma, ela pode ser implementada por qualquer estrututura
func generica(interf interface{}) {

	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	// o println recebe qualquer tipo, é um bom exemplo de interface generica

	//maior bagunça possível, nunca fazer isso
	mapa := map[interface{}]interface{}{
		1:   "string",
		"2": true,
	}

	fmt.Println(mapa)
}
