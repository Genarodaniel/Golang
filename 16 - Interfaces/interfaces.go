package main

import (
	"fmt"
	"math"
)

//interfaces são utilizadas quando precisamos ter flexibilidade entre tipos de dados

// criando a interface
// a interface diz que todas as estruturas que a implementam tem que obrigatoriamente ter seus metodos
// qualquer struct que seja criada, e tenha os métodos de alguma interface, pode implementá-la
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é: %0.2f \n", f.area())
}

func main() {

	r := retangulo{10, 20}
	//chamada do método escrever area, correta pelo fato de que r implementa o método area da interface forma
	escreverArea(r)

	c := circulo{12.2}
	escreverArea(c)

}
