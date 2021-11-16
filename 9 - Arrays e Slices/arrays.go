package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array [5]int
	array[0] = 2
	array[2] = 3
	fmt.Println(array)

	//arrays tem tipos e tamanhos fixos, e não é possível ter mais e um tipo de dado diferente dentro de um array
	array2 := [5]string{"Posição1", "Posição2", "Posição3", "Posição4", "Posição5"}
	fmt.Println(array2)

	//declarando um array com tamanho variavel
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	//SLICES

	slice := []int{10, 20, 30, 40, 50, 60, 40}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 19)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	//Arrays Internos

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 20)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
