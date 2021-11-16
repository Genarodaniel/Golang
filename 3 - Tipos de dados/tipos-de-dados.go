package main

import (
	"errors"
	"fmt"
)

func main() {

	//dividido entre 2 tipos (inteiros e reais)

	//tipos inteiros
	//int8, int16, int32, int64 diferença somente na quantidade de bits que podem receber

	var numero int64 = -444232232332313123
	//int suporta negativo
	fmt.Println(numero)

	var numero2 uint32 = 3333333
	//uint (unsigned int) não suporta negativos
	fmt.Println(numero2)

	//usado para dar o valor de apenas int(sem 32 ou 64) é usada a quantidade de bits da maquina
	var numero0 int = 321231
	fmt.Println(numero0)

	//apelido para int
	// É a mesma coisa que int32
	var numero3 rune = 12312
	fmt.Println(numero3)

	//BYTE = uint8
	//apelidos para uint
	var byte = 123
	fmt.Println(byte)

	//declaracao de numeros reais (floats)
	var numeroReal1 float32 = 123.442
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.442
	fmt.Println(numeroReal2)

	//usado para dar o valor de apenas float(sem 32 ou 64 para os tipos float) é usada a quantidade de bits da maquina
	numeroReal3 := 123123.445
	fmt.Println(numeroReal3)

	//Declaração de strings

	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	// definicao de char, que só pode conter um unico caracter e é convertido para int(rune), o char é definido com aspas simples
	char := 'B'
	fmt.Println(char)

	// fim das strings

	//BOOLEAN

	var bolleano bool = true
	fmt.Println(bolleano)
	// valor zero de bool é false

	var erro error = errors.New("erro interno")
	fmt.Println(erro)

}
