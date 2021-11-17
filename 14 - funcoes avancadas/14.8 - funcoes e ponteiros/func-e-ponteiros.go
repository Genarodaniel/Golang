package main

import "fmt"

func inverter(numero *int) {
	//* para que o valor a variavel a receber o valor, seja a referencia passada por parametro

	*numero = *numero * -1
}

func main() {

	numero := 20
	fmt.Println(numero)
	inverter(&numero) // & comercial para passar a referencia de numero como parametro
	fmt.Println("numero dps de invertido", numero)

}
