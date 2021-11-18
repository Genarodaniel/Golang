package formas

import (
	"math"
)

//interfaces são utilizadas quando precisamos ter flexibilidade entre tipos de dados

// criando a interface
// a interface diz que todas as estruturas que a implementam tem que obrigatoriamente ter seus metodos
// qualquer struct que seja criada, e tenha os métodos de alguma interface, pode implementá-la
type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.raio * c.raio)
}
