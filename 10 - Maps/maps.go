package main

import "fmt"

func main() {

	//dentro do colchete o tipo das chaves e fora dele o tipo do valor (todos os valores devem ser do mesmo tipo)
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	// para acessar os indices do map, é necessário utilizar o valor do atributo em colchetes, exemplo:
	fmt.Println(usuario["nome"])

	// Maps aninhados
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "pedro",
			"ultimo":   "silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	//deletar uma chave
	delete(usuario2, "nome")

	// adicionar uma chave

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
