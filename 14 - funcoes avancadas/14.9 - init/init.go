package main

import "fmt"

//init é uma função que será executada antes da função main
// é possível ter uma função init por arquivo
// é comunmente utilzada par ainicializar variaveis globais

func main() {

	fmt.Println("funcao main sendo executada")
}

func init() {
	fmt.Println("função executada no inicio")
}
