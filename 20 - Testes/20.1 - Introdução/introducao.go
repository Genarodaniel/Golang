package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

// go test ./... para testar todos os pacotes que estão dentro da pasta em que o comando rodou
// go test -v mostra todos os verboses do test, todas funções que foram executadas //bem melhor
// t.Parallel() mostra que o teste será rodado em paralelo com os outros testes com a mesma chamada de função
// go test --coverage para verificar quantos % estão sendo cobertos pelo teste
// o teste sempre testa os tipos de retorno esperados
// go test --coverprofile cobertura.txt mostra todos dados que não estão cobertos em um txt
// go tool cover --func=cobertura.txt para ler o arquivo txt gerado
// go tool cover --html=cobertura.txt para mostrar o html com o visual das linhas não cobertas
// para usar pacotes diferentes, utilize o _test dps do nome do pacote (endereco_test)

func main() {

	tipoDeEndereco := enderecos.TipoDeEnderecos("Avenida Paulista")
	fmt.Println(tipoDeEndereco)
}
